package api

import (
	"DEMOX_ADMINAUTH/internal/app/role/rolemodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/router/irouter"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleDel struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewRoleDel(c *gin.Context, appctx *ctx.AppContext) irouter.IHandler {
	return &RoleDel{hd.NewHd(c), c, appctx}
}

// Do
// @api 	| admin | 6 | 删除用户
// @path 	| /api/role/:id
// @method 	| DELETE
// @header 	|n Authorization |d token |e tokenstring |c 鉴权 |t string
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t type
// @response | hd.ErrResponse | 500 RESPONSE
// @tbtitle  | msg 数据
// @tbrow    |n msg |e 记录不存在
func (this *RoleDel) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}
	if err := this.Delete(id); err != nil {
		return err
	} else {
		this.RepOk()
		return nil
	}
}

func (this *RoleDel) Delete(id int64) error {
	db := this.appctx.Db
	po := &rolemodel.RolePo{}
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

	this.appctx.AuthCheck.LoadPolicy()
	return nil
}
