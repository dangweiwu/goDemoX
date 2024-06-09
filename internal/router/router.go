package router

import (
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/middler"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"github.com/gin-gonic/gin"
)

/**
路由基础定义
*/

type Router struct {
	Root *gin.RouterGroup
	Jwt  *gin.RouterGroup //jwt登陆
	Auth *gin.RouterGroup //权限
}

func NewRouter(g *gin.Engine, appctx *ctx.AppContext) *Router {
	return &Router{
		Root: g.Group("/api"),
		Jwt:  g.Group("/api", middler.TokenPase(appctx), middler.LoginCode(appctx)),
		Auth: g.Group("/api", middler.TokenPase(appctx), middler.LoginCode(appctx), middler.CheckAuth(appctx)),
	}
}

func Do(appctx *ctx.AppContext, f HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := f(c, appctx).Do(); err != nil {
			switch err.(type) {
			case hd.ErrResponse:
				//多语言用
				c.JSON(400, err)
			default:
				c.JSON(400, &hd.ErrResponse{hd.MSG, "", err.Error()})
			}
		}
	}
}
