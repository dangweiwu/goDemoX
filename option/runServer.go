package option

import (
	"DEMOX_ADMINAUTH/internal/app"
	"DEMOX_ADMINAUTH/internal/config"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/middler"
	"DEMOX_ADMINAUTH/internal/pkg/api/apiserver"
	"DEMOX_ADMINAUTH/internal/pkg/fullurl"
	"DEMOX_ADMINAUTH/internal/pkg/logx"
	"DEMOX_ADMINAUTH/internal/pkg/observe/metricx"
	"DEMOX_ADMINAUTH/internal/pkg/observe/tracex"
	"github.com/dangweiwu/ginpro/pkg/yamconfig"
	"github.com/gin-gonic/gin"
)

type RunServe struct {
	ApiHost string `long:"apihost" description:"api启动host"`
}

func (this *RunServe) Execute(args []string) error {
	//time.FixedZone()
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

	//observe 可观测性
	// trace 链路跟踪
	if c.Trace.Enable {
		tracex.InitTrace(tracex.Config{
			EndpointUrl: c.Trace.EndpointUrl,
			Auth:        c.Trace.Auth,
			ServerName:  c.Trace.ServerName,
			StreamName:  c.Trace.StreamName,
		})
		tracex.Run()
		logx.Msg("trace启动").Info(appctx.Log)
	}
	// metric
	if c.Metric.Enable {
		metricx.InitMetric(metricx.Config{
			EndpointUrl: c.Metric.EndpointUrl,
			Auth:        c.Metric.Auth,
			ServerName:  c.Metric.ServerName,
			StreamName:  c.Metric.StreamName,
		})
		metricx.Run()
		logx.Msg("metric启动").Info(appctx.Log)
	}

	//中间件 全局
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

	//可观测性结束
	tracex.Stop()
	metricx.Stop()

	//结束
	appctx.Close()
	return nil
}
