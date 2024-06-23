package api

import (
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/router"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminDel struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewAdminDel(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &AdminDel{hd.NewHd(c), c, appctx}
}

// Do
// @api | admin | 5 | 删除用户
// @path | /api/admin/:id
// @method | DELETE
// @header 	|n Authorization |d token |e tokenstring |c 鉴权 |t string
// @urlparam |n id |d 用户ID |v required |t int    |e 1
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t type
// @response | hd.ErrResponse | 500 RESPONSE
// @tbtitle  | msg 数据
// @tbrow    |n msg |e 禁止删除自己
// @tbrow    |n msg |e 记录不存在
func (this *AdminDel) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}

	uid, err := this.appctx.GetUid(this.ctx)
	if err != nil {
		return err
	}

	if id == uid {
		return errors.New("禁止删除自己")
	}

	if err := this.Delete(id); err != nil {
		return err
	} else {
		this.RepOk()
		return nil
	}
}

func (this *AdminDel) Delete(id int64) error {
	db := this.appctx.Db
	po := &adminmodel.AdminPo{}
	po.ID = id
	if r := db.Take(po); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}
	if r := db.Delete(po); r.Error != nil {
		return r.Error
	}
	return nil
}
