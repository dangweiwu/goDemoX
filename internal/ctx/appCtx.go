package ctx

import (
	"github.com/go-redis/redis/v8"
	errs "github.com/pkg/errors"
	"go.uber.org/zap"
	"goDemoX/internal/app/auth/authcheck"
	"goDemoX/internal/config"
	"goDemoX/internal/pkg/db/mysqlx"
	"goDemoX/internal/pkg/db/redisx"
	"goDemoX/internal/pkg/logx"
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
	SelfCtxI
}

func NewAppContext(c config.Config) (*AppContext, error) {
	//初始化日志
	appctx := &AppContext{}
	appctx.StartTime = time.Now()
	appctx.Config = c
	appctx.SelfCtxI = NewSelfCtx(appctx)

	if lg, err := logx.New(c.Log); err != nil {
		return nil, err
	} else {
		appctx.Log = lg
	}

	//初始化数据库
	if db, err := mysqlx.NewDb(c.Mysql); err != nil {
		return nil, errs.WithMessage(err, "err init db")
	} else {
		appctx.Db = db.GetDb()
		logx.Msg("数据库链接成功").Info(appctx.Log)
	}

	//初始化redis
	if redisCli, err := redisx.NewRedis(c.Redis); err != nil {
		return nil, errs.WithMessage(err, "err init redis")
	} else {
		appctx.Redis = redisCli.GetDb()
		logx.Msg("redis链接成功").Info(appctx.Log)
	}

	//初始化权限
	if ck, err := authcheck.NewAuthCheckout(appctx.Db); err != nil {
		return nil, err
	} else {
		appctx.AuthCheck = ck
		logx.Msg("casbin初始化完毕").Info(appctx.Log)

	}

	return appctx, nil
}
