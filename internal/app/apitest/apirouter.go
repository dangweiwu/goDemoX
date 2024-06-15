package apitest

import (
	"DEMOX_ADMINAUTH/internal/app/apitest/api"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/middler/observe"
	"DEMOX_ADMINAUTH/internal/router"
)

// @group | apitest | 20 | test | 测试api
func Route(r *router.Router, appctx *ctx.AppContext) {

	r.Root.GET("/test",
		observe.RequestTotal("api_total"),
		observe.RequestDuration("api_duration"),
		observe.Trace("test"),
		router.Do(appctx, api.NewApiTest))

}
