package api

import (
	"DEMOX_ADMINAUTH/internal/app/my/mymodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/pkg/jwtx"
	"DEMOX_ADMINAUTH/internal/router"
	"errors"
	"github.com/gin-gonic/gin"
	errs "github.com/pkg/errors"
	"gorm.io/gorm"
)

/*
修改我的信息
*/

type MyUpdate struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewMyUpdate(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &MyUpdate{hd.NewHd(c), c, appctx}
}

// Do
// @api     | me | 4 | 修改我的信息
// @path 	| /api/my
// @method 	| PUT
// @header   |n Authorization |d token  |t string |c 鉴权
// @form     | mymodel.MyForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *MyUpdate) Do() error {
	var err error
	uid, err := jwtx.GetUid(this.ctx)

	po := &mymodel.MyForm{}

	err = this.Bind(po)
	if err != nil {
		return err
	}
	po.ID = uid
	err = this.Update(po)
	if err != nil {
		return err
	}
	this.RepOk()
	return nil
}

func (this *MyUpdate) Update(rawpo *mymodel.MyForm) error {
	po := &mymodel.MyForm{}
	//校验
	if r := this.appctx.Db.Model(po).Where("id=?", rawpo.ID).Take(po); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}
	/*
		if err := this.valid(rawpo); err != nil {
			return err
		}
	*/
	//更新
	if r := this.appctx.Db.Model(rawpo).Select("phone", "name", "memo", "email").Updates(rawpo); r.Error != nil {
		return r.Error
	}
	return nil

}

func (this *MyUpdate) valid(po *mymodel.MyForm) error {
	var ct = int64(0)

	if po.Phone != "" {
		if r := this.appctx.Db.Model(po).Where("account = ?", po.Phone).Count(&ct); r.Error != nil {
			return errs.WithMessage(r.Error, "校验失败")
		} else if ct != 0 {
			return errors.New("手机号已存在")
		}
	}

	if po.Email != "" {
		if r := this.appctx.Db.Model(po).Where("email = ?", po.Email).Count(&ct); r.Error != nil {
			return errs.WithMessage(r.Error, "校验失败")
		} else if ct != 0 {
			return errors.New("Email已存在")
		}
	}
	return nil
}
