package admin

import (
	"DEMOX_ADMINAUTH/internal/app/admin/api"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/middler/observe"
	"DEMOX_ADMINAUTH/internal/router"
)

// @group | admin | 1 | 用户管理 | 系统用户管理 增删查改
func Route(r *router.Router, appctx *ctx.AppContext) {
	r.Auth.GET("/admin", observe.Trace("getAdmin"), router.Do(appctx, api.NewAdminQuery))

	r.Auth.POST("/admin", observe.Trace("createAdmin"), router.Do(appctx, api.NewAdminCreate))

	r.Auth.PUT("/admin/:id", router.Do(appctx, api.NewAdminUpdate))

	r.Auth.DELETE("/admin/:id", router.Do(appctx, api.NewAdminDel))

	r.Auth.PUT("/admin/resetpwd/:id", router.Do(appctx, api.NewResetPassword))
}
