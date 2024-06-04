package app

import (
	"DEMOX_ADMINAUTH/internal/app/admin"
	"DEMOX_ADMINAUTH/internal/app/auth"
	"DEMOX_ADMINAUTH/internal/app/my"
	"DEMOX_ADMINAUTH/internal/app/role"
	"DEMOX_ADMINAUTH/internal/app/sys"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/router"
	"github.com/gin-gonic/gin"
)

var routes = []func(r *router.Router, appctx *ctx.AppContext){
	admin.Route,
	my.Route,
	auth.Route,
	role.Route,
	sys.Route,
}

func RegisterRoute(engine *gin.Engine, appctx *ctx.AppContext) {
	r := router.NewRouter(engine, appctx)
	//regroute
	for _, v := range routes {
		v(r, appctx)
	}
}
