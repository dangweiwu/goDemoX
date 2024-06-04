package api

import (
	"DEMOX_ADMINAUTH/internal/app/my/mymodel"
	"DEMOX_ADMINAUTH/internal/app/role/rolemodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/router/irouter"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleUpdate struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewRoleUpdate(c *gin.Context, appctx *ctx.AppContext) irouter.IHandler {
	return &RoleUpdate{hd.NewHd(c), c, appctx}
}

// Do
// @api 	| role | 2 | 修改角色
// @path 	| /api/role/:id
// @method 	| PUT
// @urlparam |n id |d 角色ID   |v required |t int    |e 1
// @header   |n Authorization |d token  |t string |c 鉴权
// @form     | rolemodel.RoleUpdate
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *RoleUpdate) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}
	po := &rolemodel.RoleUpdate{}
	err = this.Bind(po)
	if err != nil {
		return err
	}
	po.ID = id
	err = this.Update(po)
	if err != nil {
		return err
	}
	this.RepOk()
	return nil
}

func (this *RoleUpdate) Update(po *rolemodel.RoleUpdate) error {
	db := this.appctx.Db
	tmpPo := &rolemodel.RolePo{}
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

	if tmpPo.Status != po.Status {
		this.appctx.AuthCheck.LoadPolicy()
		//删除缓存状态
		this.appctx.Redis.Del(context.Background(), mymodel.ROLE_STATUS+tmpPo.Code)
	}

	return nil
}
