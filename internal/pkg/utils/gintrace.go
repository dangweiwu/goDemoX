package utils

import (
	"github.com/gin-gonic/gin"
	"goDemoX/internal/pkg/observe/tracex"
)

func WithGinTraceStart(c *gin.Context, name string) tracex.Spanx {
	ctx, span := tracex.Start(c.Request.Context(), name)
	c.Request = c.Request.WithContext(ctx)
	return span
}
