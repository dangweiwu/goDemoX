package ctx

import (
	"github.com/gin-gonic/gin"
	"goDemoX/internal/pkg/jwtx"
)

type SelfCtxI interface {
	Close()
	GetUid(ctx *gin.Context) (int64, error)
	GetRole(ctx *gin.Context) (string, error)
}

type SelfCtx struct {
	*AppContext
}

func NewSelfCtx(ctx *AppContext) SelfCtxI {
	return &SelfCtx{
		AppContext: ctx,
	}
}

func (this *SelfCtx) Close() {
	this.Redis.Close()
}

func (this *SelfCtx) GetUid(ctx *gin.Context) (int64, error) {
	return jwtx.GetUid(ctx)
}

func (this *SelfCtx) GetRole(ctx *gin.Context) (string, error) {

	return jwtx.GetRole(ctx)
}
