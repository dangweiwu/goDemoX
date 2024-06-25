package router

import (
	"github.com/gin-gonic/gin"
	"goDemoX/internal/ctx"
)

type (
	IHandler interface {
		Do() error
	}
	HandlerFunc func(c *gin.Context, appctx *ctx.AppContext) IHandler
)

type RegRoute func(r *Router, appctx *ctx.AppContext)
