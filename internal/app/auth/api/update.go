package api

import (
	"DEMOX_ADMINAUTH/internal/app/auth/authmodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/router/irouter"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthUpdate struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewAuthUpdate(c *gin.Context, appctx *ctx.AppContext) irouter.IHandler {
	return &AuthUpdate{hd.NewHd(c), c, appctx}
}

// Do
// @api 	| auth | 2 | 修改权限
// @path 	| /api/auth/:id
// @method 	| PUT
// @urlparam |n id |d 权限ID   |v required |t int    |e 1
// @header   |n Authorization |d token  |t string |c 鉴权
// @form     | authmodel.AuthUpdateForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *AuthUpdate) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}
	po := &authmodel.AuthUpdateForm{}
	err = this.Bind(po)
	if err != nil {
		return err
	}
	po.ID = id
	err = this.Update(po)
	if err != nil {
		return err
	}

	this.appctx.AuthCheck.LoadPolicy()
	this.RepOk()
	return nil
}

func (this *AuthUpdate) Update(po *authmodel.AuthUpdateForm) error {
	db := this.appctx.Db
	tmpPo := &authmodel.AuthPo{}
	tmpPo.ID = po.ID
	if r := db.Model(tmpPo).Take(tmpPo); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}
	//更新
	if r := db.Updates(po); r.Error != nil {
		return r.Error
	}
	return nil
}
