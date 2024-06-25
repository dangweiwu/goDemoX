package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goDemoX/internal/app/my/mymodel"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg/api/hd"
	"goDemoX/internal/router"
	"gorm.io/gorm"
)

/*
获取我的信息
*/
type MyInfo struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewMyInfo(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &MyInfo{hd.NewHd(c), c, appctx}
}

// Do
// @api | me | 3 | 我的详情
// @path | /api/my
// @method | GET
// @header |n Authorization |d token |t string |c 鉴权
// @response | mymodel.MyInfo | 200 Response
func (this *MyInfo) Do() error {

	uid, err := this.appctx.GetUid(this.ctx)
	if err != nil {
		return err
	}

	po := &mymodel.MyInfo{}
	if r := this.appctx.Db.Model(po).Where("id = ?", uid).Take(po); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}
	this.Rep(po)
	return nil
}
