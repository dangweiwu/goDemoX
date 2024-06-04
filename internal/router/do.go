package router

import (
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/router/irouter"
	"github.com/gin-gonic/gin"
)

func Do(appctx *ctx.AppContext, f irouter.HandlerFunc) func(c *gin.Context) {
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
