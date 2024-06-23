package testapp

import (
	"github.com/gin-gonic/gin"
)

type TestSelfCtx struct {
	*TestApp
}

func NewTestSelfCtx(c *TestApp) *TestSelfCtx {
	return &TestSelfCtx{c}
}

func (this *TestSelfCtx) Close() {
	this.dbEngine.Close()
	this.rdbEngine.Close()
}
func (this *TestSelfCtx) GetUid(ctx *gin.Context) (int64, error) {
	return this.TestApp.GetUid(ctx)
}

func (this *TestSelfCtx) GetRole(ctx *gin.Context) (string, error) {

	return this.TestApp.GetRole(ctx)
}
