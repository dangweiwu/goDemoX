package api

import (
	"DEMOX_ADMINAUTH/internal/app/sys/sysmodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/pkg/observe/metricx"
	"DEMOX_ADMINAUTH/internal/pkg/observe/tracex"
	"DEMOX_ADMINAUTH/internal/router"
	"github.com/gin-gonic/gin"
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
