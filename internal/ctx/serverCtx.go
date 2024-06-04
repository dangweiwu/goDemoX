package ctx

import (
	"DEMOX_ADMINAUTH/internal/app/auth/authcheck"
	"DEMOX_ADMINAUTH/internal/config"
	"DEMOX_ADMINAUTH/internal/pkg/log"
	"github.com/dangweiwu/ginpro/pkg/mysqlx"
	"github.com/dangweiwu/ginpro/pkg/redisx"
	"github.com/dangweiwu/ginpro/pkg/tracex"
	"github.com/go-redis/redis/v8"
	errs "github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

// 所有资源放在此处
type AppContext struct {
	StartTime time.Time
	Config    config.Config
	Log       *zap.Logger
	Db        *gorm.DB
	Redis     *redis.Client
	AuthCheck *authcheck.AuthCheck
	Tracer    *tracex.Tracex
}

func NewAppContext(c config.Config) (*AppContext, error) {
	//初始化日志
	appctx := &AppContext{}
	appctx.StartTime = time.Now()
	appctx.Tracer = tracex.NewTrace(c.App.Name)
	appctx.Config = c

	if lg, err := log.New(c.Log); err != nil {
		return nil, err
	} else {
		appctx.Log = lg
	}

	//初始化数据库
	db := mysqlx.NewDb(c.Mysql)
	if d, err := db.GetDb(); err != nil {
		return nil, errs.WithMessage(err, "err init db")
	} else {
		d.Debug()
		appctx.Db = d
		log.Msg("数据库链接成功").Info(appctx.Log)
	}

	//初始化redis
	if redisCli, err := redisx.NewRedis(c.Redis).GetDb(); err != nil {
		return nil, errs.WithMessage(err, "err init redis")
	} else {
		appctx.Redis = redisCli
		log.Msg("redis链接成功").Info(appctx.Log)
	}

	//初始化权限
	if ck, err := authcheck.NewAuthCheckout(appctx.Db); err != nil {
		return nil, err
	} else {
		appctx.AuthCheck = ck
		log.Msg("casbin初始化完毕").Info(appctx.Log)

	}

	//追踪启动
	//if c.Trace.Enable {
	//	appctx.Tracer.SetEnable(true)
	//} else {
	//	appctx.Tracer.SetEnable(false)
	//}
	//
	////指标采集启动
	//if c.Prom.Enable {
	//	metric.SetEnable(true)
	//} else {
	//	metric.SetEnable(false)
	//}

	return appctx, nil
}
