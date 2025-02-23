package role

import (
	"goDemoX/internal/app/role/api"
	"goDemoX/internal/ctx"
	"goDemoX/internal/router"
)

// @group | role | 5 | 角色管理
func Route(r *router.Router, appctx *ctx.AppContext) {

	r.Auth.GET("/role", router.Do(appctx, api.NewRoleQuery))

	r.Auth.POST("/role", router.Do(appctx, api.NewRoleCreate))

	r.Auth.PUT("/role/:id", router.Do(appctx, api.NewRoleUpdate))

	r.Auth.DELETE("/role/:id", router.Do(appctx, api.NewRoleDel))

	r.Auth.PUT("/role/auth/:id", router.Do(appctx, api.NewSetAuth))

}
