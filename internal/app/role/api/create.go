package api

import (
	"DEMOX_ADMINAUTH/internal/app/role/rolemodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/router"
	"errors"
	"github.com/gin-gonic/gin"
)

type RoleCreate struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewRoleCreate(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &RoleCreate{hd.NewHd(c), c, appctx}
}

// Do
// @api | role | 1 | 创建角色
// @path    | /api/role
// @method  | POST
// @header  |n Authorization |d token |t string |c 鉴权
// @form    | rolemodel.RoleForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *RoleCreate) Do() error {
	//数据源
	po := &rolemodel.RoleForm{}
	err := this.Bind(po)
	if err != nil {
		return err
	}

	err = this.Create(po)
	if err != nil {
		return err
	}
	this.RepOk()
	return nil
}

func (this *RoleCreate) Create(po *rolemodel.RoleForm) error {
	db := this.appctx.Db
	//验证是否已创建 或者其他检查
	tmpPo := &rolemodel.RolePo{}
	if r := db.Model(tmpPo).Where("code = ?", po.Code).Take(tmpPo); r.Error == nil {
		return errors.New("角色编码已存在")
	}

	if r := db.Create(po); r.Error != nil {
		return r.Error
	}
	return nil
}
