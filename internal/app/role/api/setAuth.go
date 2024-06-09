package api

import (
	"DEMOX_ADMINAUTH/internal/app/my/mymodel"
	"DEMOX_ADMINAUTH/internal/app/role/rolemodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/router"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SetAuth struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewSetAuth(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &SetAuth{hd.NewHd(c), c, appctx}
}

// Do
// @api 	| role | 3 | 修改角色
// @path 	| /role/auth/:id
// @method 	| PUT
// @urlparam |n id |d 角色ID   |v required |t int    |e 1
// @header   |n Authorization |d token  |t string |c 鉴权
// @form     | rolemodel.RoleAuthForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *SetAuth) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}

	po := &rolemodel.RoleAuthForm{}
	err = this.Bind(po)
	if err != nil {
		return err
	}
	po.ID = id
	if err := this.Update(po); err != nil {
		return err
	}
	this.RepOk()
	return nil
}

func (this *SetAuth) Update(po *rolemodel.RoleAuthForm) error {
	db := this.appctx.Db
	tmpPo := &rolemodel.RolePo{}
	if r := db.Model(tmpPo).Take(tmpPo); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}
	//校验是否存在

	//更新
	if r := db.Updates(po); r.Error != nil {
		return r.Error
	}
	//删除
	this.appctx.Redis.Del(context.Background(), mymodel.ROLE_AUTH+tmpPo.Code)
	this.appctx.AuthCheck.LoadPolicy()
	return nil
}
