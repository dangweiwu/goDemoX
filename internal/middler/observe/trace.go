package observe

import (
	"DEMOX_ADMINAUTH/internal/pkg/observe/tracex"
	"DEMOX_ADMINAUTH/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
)

func Trace(tracename string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !tracex.GetEable() {
			c.Next()
			return
		}

		//spanName := c.FullPath()
		//if spanName == "" {
		//	spanName = fmt.Sprintf("HTTP %s route not found", c.Request.Method)
		//}

		span := utils.WithGinTraceStart(c, tracename)

		defer span.End()
		span.SetAttributes(attribute.String("path", c.Request.Method+":"+c.FullPath()))
		c.Next()

		status := c.Writer.Status()
		span.SetAttributes(attribute.Int("http.status", status))

		if len(c.Errors) > 0 {
			span.SetAttributes(attribute.String("gin.errors", c.Errors.String()))
		}
	}
}
