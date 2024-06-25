package middler

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"goDemoX/internal/ctx"
)

// 批量注册全局中间件
func RegMiddler(appctx *ctx.AppContext) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		requestid.New(),
		Recovery(appctx),
		HttpLog(appctx),
		Cors(),
	}
}
