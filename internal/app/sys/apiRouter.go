package sys

import (
	"goDemoX/internal/app/sys/api"
	"goDemoX/internal/ctx"
	"goDemoX/internal/router"
)

// @group | sys | 3 | 系统设置 | 包括链路追踪,指标采集
func Route(r *router.Router, appctx *ctx.AppContext) {
	r.Auth.GET("/sys", router.Do(appctx, api.NewSysQuery))
	r.Auth.PUT("/sys", router.Do(appctx, api.NewSysAct))
}
