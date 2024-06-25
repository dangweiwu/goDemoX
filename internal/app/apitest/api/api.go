package api

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg/api/hd"
	"goDemoX/internal/pkg/utils"
	"goDemoX/internal/router"
	"log"
)

type ApiTest struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

// Do
// @api | apitest | 3 | 我的详情
// @path | /api/test
// @method | GET
// @response | hd.Response | 200 Response
func NewApiTest(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &ApiTest{hd.NewHd(c), c, appctx}
}

func (this ApiTest) Do() error {

	span := utils.WithGinTraceStart(this.ctx, "do")
	defer span.End()

	span.SetAttributes(attribute.String("test", "测试"))
	//metric
	log.Println("this is test do")
	this.RepOk()
	return nil
}
