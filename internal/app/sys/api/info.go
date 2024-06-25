package api

import (
	"github.com/gin-gonic/gin"
	"goDemoX/internal/app/sys/sysmodel"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg/api/hd"
	"goDemoX/internal/pkg/observe/metricx"
	"goDemoX/internal/pkg/observe/tracex"
	"goDemoX/internal/router"
	"time"
)

type SysQuery struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewSysQuery(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &SysQuery{hd.NewHd(c), c, appctx}
}

// Do
// @api | sys | 1 | 运行状态
// @path | /api/sys
// @method | GET
// @header |n  Authorization |d 权限 |t type |c bascAuth base64(name:password)
// @response | sysmodel.SysVo | 200 Response
func (this *SysQuery) Do() error {
	vo := &sysmodel.SysVo{}
	vo.StartTime = this.appctx.StartTime.String()
	vo.RunTime = time.Now().Sub(this.appctx.StartTime).String()
	if tracex.GetEable() {
		vo.OpenTrace = "1"
	} else {
		vo.OpenTrace = "0"
	}

	if metricx.GetEnable() {
		vo.OpenMetric = "1"
	} else {
		vo.OpenMetric = "0"
	}

	this.Rep(vo)
	return nil
}
