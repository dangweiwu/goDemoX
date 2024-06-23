package api

import (
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/app/my/mymodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/router"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
)

/*
修改我的密码
*/

type UpdatePwd struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewMyUpdatePwd(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &UpdatePwd{hd.NewHd(c), c, appctx}
}

// Do
// @api     | me | 5 | 修改密码
// @path 	| /api/admin/my/password
// @method 	| PUT
// @header   |n Authorization |d token  |t string |c 鉴权
// @form     | mymodel.PasswordForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *UpdatePwd) Do() error {
	var err error
	uid, err := this.appctx.GetUid(this.ctx)

	po := &mymodel.PasswordForm{}

	err = this.Bind(po)
	if err != nil {
		return err
	}
	err = this.UpdatePwd(po, uid)
	if err != nil {
		return err
	}
	this.RepOk()
	return nil
}

func (this *UpdatePwd) UpdatePwd(form *mymodel.PasswordForm, id int64) error {
	po := &adminmodel.AdminPo{}
	if r := this.appctx.Db.Model(po).Where("id=?", id).Take(po); r.Error != nil {
		return r.Error
	}

	//校验旧密码是否正确
	if pkg.GetPassword(form.Password) != po.Password {
		return errors.New("原密码错误")
	}
	po.Password = pkg.GetPassword(form.NewPassword)
	r := this.appctx.Db.Model(po).Select("password").Updates(po)

	this.appctx.Redis.Del(context.Background(), mymodel.GetAdminRedisLoginId(int(po.ID)))
	return r.Error
}
