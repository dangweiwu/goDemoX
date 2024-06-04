package api_test

import (
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/testtool"
	"DEMOX_ADMINAUTH/internal/testtool/testctx"
	"testing"
)

var SerCtx *ctx.AppContext
var TestCtx *testctx.TestContext

func TestMain(m *testing.M) {
	config := testtool.NewTestConfig()
	//单元测试并发执行 防止数据库端口冲突
	config.Mysql.Host = "127.0.0.1:4310"
	//config.Mysql.LogLevel = 4
	ctx, err := testctx.NewTestContext(config)

	defer func() {
		ctx.Close()
	}()
	if err != nil {
		panic(err)
	}

	SerCtx, err = ctx.GetServerCtx()
	err = ctx.InitChecAuth()
	if err != nil {
		panic(err)
	}

	TestCtx = ctx
	if err != nil {
		panic(err)
	}
	m.Run()
}
