package middler

import (
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/log"
	"fmt"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
)

/*
es mapping kind = panic
*/

func Recovery(appctx *ctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				raw := c.Request.URL.RawQuery
				path := c.Request.URL.Path

				if raw != "" {
					path = path + "?" + raw
				}
				//httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					log.Msg("网络中断").Kind("api").Trace(requestid.Get(c)).ErrData(err.(error)).Err(appctx.Log)
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				log.Msg("系统异常").Kind("api").Trace(requestid.Get(c)).ErrData(err.(error)).Data(string(debug.Stack())).Err(appctx.Log)
				c.AbortWithStatus(http.StatusInternalServerError)
				c.String(500, fmt.Sprintf("%v", err))
			}
		}()
		c.Next()
	}
}
