package utils

import (
	"DEMOX_ADMINAUTH/internal/pkg/observe/tracex"
	"github.com/gin-gonic/gin"
)

func WithGinTraceStart(c *gin.Context, name string) tracex.Spanx {
	ctx, span := tracex.Start(c.Request.Context(), name)
	c.Request = c.Request.WithContext(ctx)
	return span
}
