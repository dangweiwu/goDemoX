package testctx

import (
	"DEMOX_ADMINAUTH/internal/app"
	"DEMOX_ADMINAUTH/internal/app/auth/authcheck"
	"DEMOX_ADMINAUTH/internal/config"
	"DEMOX_ADMINAUTH/internal/ctx"
	mredis "github.com/alicebob/miniredis/v2"
	"github.com/dangweiwu/ginpro/pkg/mysqlx"
	"github.com/dangweiwu/ginpro/pkg/mysqlx/mysqlfake"
	"github.com/dangweiwu/ginpro/pkg/redisx"
	"github.com/dangweiwu/ginpro/pkg/tracex"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/go-redis/redis/v8"
	errs "github.com/pkg/errors"
	"gorm.io/gorm"
)

type TestContext struct {
	Db       *gorm.DB
	Redis    *redis.Client
	Config   config.Config
	DbSer    *server.Server
	RedisSer *mredis.Miniredis
	ctx      *ctx.AppContext
}

func NewTestContext(cfg config.Config) (*TestContext, error) {
	a := &TestContext{}
	a.Config = cfg
	a.DbSer = mysqlfake.FakeMysql(a.Config.Mysql.Host, a.Config.Mysql.DbName)
	mr, err := mredis.Run()
	if err != nil {
		return nil, err
	}
	a.RedisSer = mr

	//初始化数据库
	db := mysqlx.NewDb(cfg.Mysql)
	if d, err := db.GetDb(); err != nil {
		return nil, errs.WithMessage(err, "err init db")
	} else {
		a.Db = d
	}

	//初始化redis
	cfg.Redis.Addr = mr.Addr()
	if redisCli, err := redisx.NewRedis(cfg.Redis).GetDb(); err != nil {
		return nil, errs.WithMessage(err, "err init redis")

	} else {
		a.Redis = redisCli
	}

	return a, nil
}

func (this *TestContext) GetServerCtx() (*ctx.AppContext, error) {
	a := &ctx.AppContext{}
	if lg, err := logx.New(this.Config.Log); err != nil {
		return nil, err
	} else {
		a.Log = lg
	}
	a.Config = this.Config
	a.Db = this.Db
	a.Redis = this.Redis
	a.Tracer = tracex.NewTrace("xx")

	//注册数据库
	app.Regdb(a)
	this.ctx = a
	return a, nil
}

func (this *TestContext) InitChecAuth() error {
	//初始化权限
	if ck, err := authcheck.NewAuthCheckout(this.Db); err != nil {
		return err
	} else {
		this.ctx.AuthCheck = ck
		logx.Msg("casbin初始化完毕").Info(this.ctx.Log)
	}
	return nil
}

func (this *TestContext) Close() {
	this.DbSer.Close()
	this.RedisSer.Close()
}
