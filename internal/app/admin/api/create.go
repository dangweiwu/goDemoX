package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	errs "github.com/pkg/errors"
	"goDemoX/internal/app/admin/adminmodel"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg"
	"goDemoX/internal/pkg/api/hd"
	"goDemoX/internal/router"
)

type AdminCreate struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewAdminCreate(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &AdminCreate{hd.NewHd(c), c, appctx}
}

// Do
// @api     | admin | 1 |创建用户
// @path    | /api/admin
// @method  | POST
// @header  |n Authorization |d token |t string |c 鉴权
// @form    | adminmodel.AdminForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *AdminCreate) Do() error {

	//数据源
	po := &adminmodel.AdminForm{}
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

func (this *AdminCreate) Create(po *adminmodel.AdminForm) error {
	db := this.appctx.Db
	//验证是否已创建 或者其他检查
	if err := this.Valid(po); err != nil {
		return err
	}

	po.Password = pkg.GetPassword(po.Password)
	if po.IsSuperAdmin == "1" {
		po.Role = ""
	}
	if r := db.Create(po); r.Error != nil {
		return r.Error
	}
	return nil
}

func (this *AdminCreate) Valid(po *adminmodel.AdminForm) error {
	db := this.appctx.Db
	var ct = int64(0)
	if r := db.Model(po).Where("account = ?", po.Account).Count(&ct); r.Error != nil {
		return errs.WithMessage(r.Error, "校验失败")
	} else if ct != 0 {
		return errors.New("账号已存在")
	}

	/*
		if po.Phone != "" {
			if r := db.Model(po).Where("phone = ?", po.Phone).Count(&ct); r.Error != nil {
				return errs.WithMessage(r.Error, "校验失败")
			} else if ct != 0 {
				return errors.New("手机号已存在")
			}
		}

		if po.Email != "" {
			if r := db.Model(po).Where("email = ?", po.Email).Count(&ct); r.Error != nil {
				return errs.WithMessage(r.Error, "校验失败")
			} else if ct != 0 {
				return errors.New("Email已存在")
			}
		}
	*/
	return nil
}
