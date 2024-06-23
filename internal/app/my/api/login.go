package api

import (
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/app/my/mymodel"
	"DEMOX_ADMINAUTH/internal/app/my/myserver"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/router"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type Login struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewLogin(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &Login{hd.NewHd(c), c, appctx}
}

// Do
// @api     | me | 1 | 登录
// @path 	| /api/login
// @method 	| POST
// @form     | mymodel.LoginForm
// @response | mymodel.LogRep |200 Response
// @response | hd.ErrResponse | 400 Response
// @tbtitle | Msg 数据
// @tbrow   |n msg |d 密码错误
// @tbrow   |n msg |d 账号不存在
// @tbrow   |n msg |d 账号被禁用
func (this *Login) Do() error {

	//数据源
	po := &mymodel.LoginForm{}
	err := this.Bind(po)
	if err != nil {
		return err
	}

	data, err := this.Login(po)
	if err != nil {
		return err
	}
	this.Rep(data)
	return nil
}
func (this *Login) Login(form *mymodel.LoginForm) (*mymodel.LogRep, error) {
	var (
		token        string
		refreshToken string
		logcode      string
		err          error
	)
	po, err := this.Valid(form)
	if err != nil {
		return nil, err
	}

	pwd := pkg.GetPassword(form.Password)
	if pwd != po.Password {
		return nil, errors.New("密码错误")
	}

	logcode, err = myserver.NewLogCode(po.ID, this.appctx.Redis)
	if err != nil {
		fmt.Println("err====redis")
		return nil, err
	}

	logopt := adminmodel.LoginOpt{
		Exp:       this.appctx.Config.Jwt.Exp,
		LoginCode: logcode,
		Secret:    this.appctx.Config.Jwt.Secret,
	}
	token, refreshToken, err = po.Login(&logopt)
	if err != nil {
		return nil, err
	}

	return &mymodel.LogRep{
		token,
		time.Now().Unix() + logopt.Exp - 60,
		refreshToken,
	}, nil

}

func (this *Login) Valid(form *mymodel.LoginForm) (*adminmodel.AdminPo, error) {
	po := &adminmodel.AdminPo{}
	if r := this.appctx.Db.Model(po).Where("account=?", form.Account).Take(po); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return po, errors.New("账号不存在")
		} else {
			return po, r.Error
		}
	}

	if pkg.GetPassword(form.Password) != po.Password {
		return nil, errors.New("密码错误")
	}

	if po.Status == "0" {
		return nil, errors.New("账号被禁用")
	}

	return po, nil
}
