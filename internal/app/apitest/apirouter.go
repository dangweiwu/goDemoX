package apitest

import (
	"goDemoX/internal/app/apitest/api"
	"goDemoX/internal/ctx"
	"goDemoX/internal/middler/observe"
	"goDemoX/internal/router"
)

// @group | apitest | 20 | test | 测试api
func Route(r *router.Router, appctx *ctx.AppContext) {
	r.Root.GET("/test",
		observe.RequestTotal("api_total"),
		observe.RequestDuration("api_duration"),
		observe.Trace("test"),
		router.Do(appctx, api.NewApiTest))
}
