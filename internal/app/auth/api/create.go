package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goDemoX/internal/app/auth/authmodel"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg/api/hd"
	"goDemoX/internal/router"
)

type AuthCreate struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewAuthCreate(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &AuthCreate{hd.NewHd(c), c, appctx}
}

// Do
// @api | auth | 1 | 创建权限
// @path    | /api/auth
// @method  | POST
// @header  |n Authorization |d token |t string |c 鉴权
// @form    | authmodel.AuthForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *AuthCreate) Do() error {
	//数据源
	po := &authmodel.AuthForm{}
	err := this.Bind(po)
	if err != nil {
		return err
	}

	err = this.Create(po)
	if err != nil {
		return err
	}

	this.appctx.AuthCheck.LoadPolicy()
	this.RepOk()
	return nil
}

func (this *AuthCreate) Create(po *authmodel.AuthForm) error {
	db := this.appctx.Db
	//验证是否已创建 或者其他检查
	tmpPo := &authmodel.AuthPo{}
	if r := db.Model(po).Where("code = ?", po.Code).Take(tmpPo); r.Error == nil {
		return errors.New("权限编码已存在")
	}

	if r := db.Create(po); r.Error != nil {
		return r.Error
	}
	return nil
}
