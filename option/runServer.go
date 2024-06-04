package option

import (
	"DEMOX_ADMINAUTH/internal/app"
	"DEMOX_ADMINAUTH/internal/config"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/middler"
	"DEMOX_ADMINAUTH/internal/pkg/api/apiserver"
	"DEMOX_ADMINAUTH/internal/pkg/fullurl"
	"DEMOX_ADMINAUTH/internal/pkg/log"
	"DEMOX_ADMINAUTH/internal/pkg/observe/tracex"
	"github.com/dangweiwu/ginpro/pkg/yamconfig"
	"github.com/gin-gonic/gin"
)

type RunServe struct {
	ApiHost string `long:"apihost" description:"api启动host"`
}

func (this *RunServe) Execute(args []string) error {
	//配置参数
	var c config.Config
	yamconfig.MustLoad(Opt.ConfigPath, &c)

	if Opt.RunServe.ApiHost != "" {
		c.Api.Host = Opt.RunServe.ApiHost
	}
	//资源初始化
	appctx, err := ctx.NewAppContext(c)
	if err != nil {
		panic(err)
	}

	//服务 中间件
	//engine := gin.Default()

	engine := gin.New()
	//trace
	//if c.Trace.Enable {
	//	tp := tracex.InitTracerHTTP(c.Trace.Endpoint, c.Trace.UrlPath, c.Trace.Auth, c.App.Name)
	//	defer func() {
	//		if err := tp.Shutdown(context.Background()); err != nil {
	//			appctx.Log.Error("Error shutting down tracer provider", zap.Error(err))
	//		}
	//	}()
	//}

	//启动promagent
	//engine.GET("/metrics", gin.BasicAuth(gin.Accounts{c.Prom.UserName: c.Prom.Password}), gin.WrapH(promhttp.Handler()))

	//observe 可观测性
	// trace 链路跟踪
	if c.Trace.Enable {
		tracex.InitTrace(tracex.Config{
			EndpointUrl: c.Trace.EndpointUrl,
			Auth:        c.Trace.Auth,
			ServerName:  c.Trace.ServerName,
			StreamName:  c.Trace.StreamName,
		})
		println(c.Trace.Auth, c.Trace.EndpointUrl, c.Trace.StreamName)
		tracex.Run()
		log.Msg("trace启动").Info(appctx.Log)
	}

	//中间件
	apiserver.RegMiddler(engine,
		apiserver.WithStatic("/view", c.Api.ViewDir),
		apiserver.WithMiddle(middler.RegMiddler(appctx)...),
	)

	//注册路由
	app.RegisterRoute(engine, appctx)

	//记录路由
	fullurl.NewFullUrl().InitUrl(engine)

	//启动
	apiserver.Run(engine, appctx.Log, c.Api)
	return nil
}
