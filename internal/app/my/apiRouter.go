package my

import (
	"DEMOX_ADMINAUTH/internal/app/my/api"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/middler/observe"
	"DEMOX_ADMINAUTH/internal/router"
)

// @group | me | 2 | 系统我的 | 包括基本信息获取修改 登录登出 token刷新
func Route(r *router.Router, appctx *ctx.AppContext) {

	r.Jwt.GET("/my", router.Do(appctx, api.NewMyInfo))

	r.Jwt.PUT("/my", router.Do(appctx, api.NewMyUpdate))

	r.Jwt.PUT("/my/password", router.Do(appctx, api.NewMyUpdatePwd))

	r.Root.POST("/login", observe.Trace("/login"), router.Do(appctx, api.NewLogin))

	r.Jwt.POST("/logout", router.Do(appctx, api.NewLogOut))

	r.Jwt.POST("/token/refresh", router.Do(appctx, api.NewRefreshToken))

	r.Jwt.GET("/my-auth", router.Do(appctx, api.NewMyAuth))
}
