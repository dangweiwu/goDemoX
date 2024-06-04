package irouter

import (
	"DEMOX_ADMINAUTH/internal/ctx"
	"github.com/gin-gonic/gin"
)

type (
	IHandler interface {
		Do() error
	}
	HandlerFunc func(c *gin.Context, appctx *ctx.AppContext) IHandler
)
