package api_test

import (
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/testtool"
	"DEMOX_ADMINAUTH/internal/testtool/testctx"
	"github.com/dangweiwu/ginpro/pkg/tracex"
	"testing"
)

var SerCtx *ctx.AppContext
var TestCtx *testctx.TestContext
var (
	Name     = "name"
	Password = "123456"
)

func TestMain(m *testing.M) {
	config := testtool.NewTestConfig()
	config.App.Name = Name
	config.App.Password = Password
	config.Trace.Enable = true
	config.Prom.Enable = true
	//单元测试并发执行 防止数据库端口冲突
	config.Mysql.Host = "127.0.0.1:4311"
	//config.Mysql.LogLevel = 4
	ctx, err := testctx.NewTestContext(config)

	defer func() {
		ctx.Close()
	}()
	if err != nil {
		panic(err)
	}

	SerCtx, err = ctx.GetServerCtx()
	SerCtx.Tracer = tracex.NewTrace("test")
	TestCtx = ctx
	if err != nil {
		panic(err)
	}
	m.Run()
}
