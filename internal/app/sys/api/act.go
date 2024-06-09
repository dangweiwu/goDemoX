package api

import (
	"DEMOX_ADMINAUTH/internal/app/sys/sysmodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/router"
	"errors"
	"github.com/dangweiwu/ginpro/pkg/metric"
	"github.com/gin-gonic/gin"
)

type SysAct struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewSysAct(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &SysAct{hd.NewHd(c), c, appctx}
}

// Do
// @api |sys| 2 | 设定开关
// @path | /api/sys
// @method | PUT
// @header |n  Authorization |d 权限 |t type |c bascAuth base64(name:password)
// @form | sysmodel.SysActForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *SysAct) Do() error {
	form := &sysmodel.SysActForm{}
	err := this.Bind(form)
	if err != nil {
		return err
	}

	name := form.Name
	if name == "" {
		return errors.New("缺少名称")
	}

	switch name {
	case "trace":
		if this.appctx.Config.Trace.Enable {
			if form.Act == "0" {
				this.appctx.Tracer.SetEnable(false)
			} else if form.Act == "1" {
				this.appctx.Tracer.SetEnable(true)
			} else {
				return errors.New("未知指令")
			}
		} else {
			return errors.New("trace 未启动")
		}
	case "metric":
		if this.appctx.Config.Prom.Enable {
			if form.Act == "0" {
				metric.SetEnable(false)
			} else if form.Act == "1" {
				metric.SetEnable(true)
			} else {
				return errors.New("未知指令")
			}
		} else {
			return errors.New("metric 未启动")
		}
	default:
		return errors.New("无效名称")
	}
	this.RepOk()
	return nil
}
