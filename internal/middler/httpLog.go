package middler

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg/jwtx"
	"goDemoX/internal/pkg/logx"
	"time"
)

var skipPath = map[string]struct{}{}
var skipMethod = map[string]struct{}{
	"GET": {},
}

func HttpLog(appctx *ctx.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		// Process request
		c.Next()

		// Log only when path is not being skipped
		_, ok := skipPath[c.FullPath()]
		_, ok2 := skipMethod[c.Request.Method]
		if !ok && !ok2 {
			// Stop timer
			TimeStamp := time.Now()
			Latency := TimeStamp.Sub(start)

			if raw != "" {
				path = path + "?" + raw
			}
			uid, _ := jwtx.GetUid(c)

			l := logx.Msg("request").Kind("api").Trace(requestid.Get(c)).FmtData("useid:%d method:%s path:%s status:%d size:%d latency:%d",
				uid, c.Request.Method, c.Request.URL.Path, c.Writer.Status(), c.Writer.Size(), int(Latency.Milliseconds()))

			if len(c.Errors) != 0 {
				l.ErrData(c.Errors[0]).Err(appctx.Log)
			}
			l.Debug(appctx.Log)
		}
	}
}
