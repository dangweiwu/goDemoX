package router

import (
	"github.com/gin-gonic/gin"
	"goDemoX/internal/ctx"
	"goDemoX/internal/middler"
)

func NewTestRouter(g *gin.Engine, appctx *ctx.AppContext) *Router {
	return &Router{
		Root: g.Group("/api"),
		Jwt:  g.Group("/api", middler.TokenPase(appctx), middler.LoginCode(appctx)),
		Auth: g.Group("/api", middler.TokenPase(appctx), middler.LoginCode(appctx)), //不进行权限验证
	}
}

func NewTestBaseRouter(g *gin.Engine, appctx *ctx.AppContext) *Router {
	return &Router{
		Root: g.Group("/api"),
		Jwt:  g.Group("/api"),
		Auth: g.Group("/api"),
	}
}
