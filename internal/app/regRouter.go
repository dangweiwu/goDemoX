package app

import (
	"github.com/gin-gonic/gin"
	"goDemoX/internal/app/admin"
	"goDemoX/internal/app/apitest"
	"goDemoX/internal/app/auth"
	"goDemoX/internal/app/my"
	"goDemoX/internal/app/role"
	"goDemoX/internal/app/sys"
	"goDemoX/internal/ctx"
	"goDemoX/internal/router"
)

var routes = []func(r *router.Router, appctx *ctx.AppContext){
	admin.Route,
	my.Route,
	auth.Route,
	role.Route,
	sys.Route,
	apitest.Route,
}

func RegisterRoute(engine *gin.Engine, appctx *ctx.AppContext) {
	r := router.NewRouter(engine, appctx)
	//regroute
	for _, v := range routes {
		v(r, appctx)
	}
}
