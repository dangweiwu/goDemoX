package middler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg/api/hd"
	"goDemoX/internal/pkg/jwtx"
)

func NoAuthErrResponse(c *gin.Context, data string) {
	c.JSON(403, hd.ErrMsg(data, "缺少权限"))
	c.Abort()
}

// 权限中间件
func CheckAuth(appctx *ctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		yes, err := jwtx.GetIsSuper(context)
		if err != nil {
			NoAuthErrResponse(context, err.Error())
			return
		}
		if yes {
			context.Next()
			return
		}
		role, err := jwtx.GetRole(context)
		if err != nil {
			NoAuthErrResponse(context, err.Error())
			return
		}

		if ok, err := appctx.AuthCheck.Check(role, context.FullPath(), context.Request.Method); ok {
			context.Next()
			return
		} else if err != nil {
			NoAuthErrResponse(context, err.Error())
		} else {
			NoAuthErrResponse(context, fmt.Sprintf("%s:%s", context.Request.Method, context.FullPath()))
		}
	}
}
