package auth

import (
	"DEMOX_ADMINAUTH/internal/app/auth/api"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/router"
)

// @group | auth | 4 | 权限管理
func Route(r *router.Router, appctx *ctx.AppContext) {

	r.Auth.GET("/auth", router.Do(appctx, api.NewAuthQuery))

	r.Auth.POST("/auth", router.Do(appctx, api.NewAuthCreate))

	r.Auth.PUT("/auth/:id", router.Do(appctx, api.NewAuthUpdate))

	r.Auth.DELETE("/auth/:id", router.Do(appctx, api.NewAuthDel))

	r.Root.GET("/allurl", router.Do(appctx, api.NewGetFullUrl))
}
