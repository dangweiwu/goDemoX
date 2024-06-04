package sys

import (
	"DEMOX_ADMINAUTH/internal/app/sys/api"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/router"
	"github.com/gin-gonic/gin"
)

// @group | sys | 3 | 系统设置 | 包括链路追踪,指标采集
func Route(r *router.Router, appctx *ctx.AppContext) {
	r.Root.GET("/sys", gin.BasicAuth(gin.Accounts{appctx.Config.App.Name: appctx.Config.App.Password}), router.Do(appctx, api.NewSysQuery))
	r.Root.PUT("/sys", gin.BasicAuth(gin.Accounts{appctx.Config.App.Name: appctx.Config.App.Password}), router.Do(appctx, api.NewSysAct))
}
