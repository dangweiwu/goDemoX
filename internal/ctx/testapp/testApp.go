package testapp

import (
	"DEMOX_ADMINAUTH/internal/app/auth/authcheck"
	"DEMOX_ADMINAUTH/internal/config"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/db/mysqlx"
	"DEMOX_ADMINAUTH/internal/pkg/db/redisx"
	"DEMOX_ADMINAUTH/internal/pkg/jwtx/jwtconfig"
	"DEMOX_ADMINAUTH/internal/pkg/logx"
	"DEMOX_ADMINAUTH/internal/pkg/observe/metricx"
	"DEMOX_ADMINAUTH/internal/pkg/observe/tracex"
	"DEMOX_ADMINAUTH/internal/pkg/utils"
	"fmt"
	mredis "github.com/alicebob/miniredis/v2"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/gin-gonic/gin"
	errs "github.com/pkg/errors"
	"time"
)

type TestApp struct {
	*ctx.AppContext
	*TestServer

	dbEngine  *server.Server
	rdbEngine *mredis.Miniredis
	GetUid    func(ctx *gin.Context) (int64, error)
	GetRole   func(ctx *gin.Context) (string, error)
}

// 初始化数据库 redis httpserver
func NewTestApp() (*TestApp, error) {

	var (
		err error
	)
	c := NewTestConfig()
	a := &TestApp{}
	a.AppContext = &ctx.AppContext{}
	a.AppContext.SelfCtxI = NewTestSelfCtx(a)
	a.StartTime = time.Now()
	a.Config = c
	a.Log, err = logx.New(c.Log)
	if err != nil {
		return nil, err
	}

	//注册所有数据库
	//app.Regdb(a.AppContext)
	a.InitDb()
	//初始化redis
	a.InitRedis()
	//初始化权限

	//if ck, err := authcheck.NewAuthCheckout(a.Db); err != nil {
	//	return nil, err
	//} else {
	//	a.AuthCheck = ck
	//	logx.Msg("casbin初始化完毕").Info(a.Log)
	//}

	//http test
	a.TestServer = NewTestServer()

	return a, nil
}

func (this *TestApp) InitAuthCheckout() error {
	//初始化权限
	if ck, err := authcheck.NewAuthCheckout(this.AppContext.Db); err != nil {
		return err
	} else {
		this.AppContext.AuthCheck = ck

	}
	return nil
}

func (this *TestApp) InitTrace() error {
	tracex.InitTrace(tracex.Config{
		EndpointUrl: this.Config.Trace.EndpointUrl,
		Auth:        this.Config.Trace.Auth,
		ServerName:  this.Config.Trace.ServerName,
		StreamName:  this.Config.Trace.StreamName,
	})
	return tracex.Run()
}

func (this *TestApp) InitMetric() error {
	metricx.InitMetric(metricx.Config{
		EndpointUrl: this.Config.Metric.EndpointUrl,
		Auth:        this.Config.Metric.Auth,
		ServerName:  this.Config.Metric.ServerName,
		StreamName:  this.Config.Metric.StreamName,
	})
	metricx.Run()
	return nil
}

func (this *TestApp) InitRedis() error {
	var err error
	this.rdbEngine, err = mredis.Run()
	if err != nil {
		return fmt.Errorf("fakeRedisErr %v", err)
	}

	this.Config.Redis.Addr = this.rdbEngine.Addr()

	if redisCli, err := redisx.NewRedis(this.Config.Redis); err != nil {
		return errs.WithMessage(err, "err init redis")

	} else {
		fmt.Println("[redis init]=========", this.Config.Redis.Addr)
		this.Redis = redisCli.GetDb()
	}
	return nil
}

func (this *TestApp) InitDb() error {
	var (
		err    error
		dbhost string
	)

	//随机生成数据库名字
	dbhost, this.dbEngine, err = utils.FakeDb(this.Config.Mysql.DbName, ":0")
	if err != nil {
		return fmt.Errorf("fake db error %v", err)
	}

	//初始化数据库
	this.Config.Mysql.Host = dbhost
	if db, err := mysqlx.NewDb(this.Config.Mysql); err != nil {
		return errs.WithMessage(err, "err init db")

	} else {

		this.Db = db.GetDb()
		this.Db.Debug()
		//数据清空
		logx.Msg("数据库链接成功").Info(this.Log)
	}

	return nil
}

func (this *TestApp) RegDb(tables ...interface{}) {
	this.Db.Migrator().DropTable(tables...)
	this.Db.Set("gorm:ble_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(tables...)

}

//func (this *TestContext) Close() {
//	this.dbEngine.Close()
//	this.rdbEngine.Close()
//}

func NewTestConfig() config.Config {
	c := config.Config{}
	c.Log = logx.Config{OutType: logx.CONSOLE, Level: logx.DEBUG, HasTimestamp: true, Formatter: logx.TXT}
	c.Mysql = mysqlx.Config{DbName: "test", LogLevel: 1}
	c.Redis = redisx.Config{Addr: ":0"}
	c.Jwt = jwtconfig.JwtConfig{
		Secret: "test",
		Exp:    3600,
	}
	return c
}
