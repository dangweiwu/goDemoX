package api

import (
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/app/my/mymodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/pkg/jwtx"
	"DEMOX_ADMINAUTH/internal/router/irouter"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Login struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewLogin(c *gin.Context, appctx *ctx.AppContext) irouter.IHandler {
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
func (this *Login) Login(form *mymodel.LoginForm) (interface{}, error) {
	po, err := this.Valid(form)
	if err != nil {
		return nil, err
	}

	pwd := pkg.GetPassword(form.Password)
	if pwd != po.Password {
		return nil, errors.New("密码错误")
	}

	//同时只能有一个jwt登陆，可拓展踢人功能
	logincode, err := this.newLoginCode(po.ID)
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	if token, err := jwtx.GenToken(
		this.appctx.Config.Jwt.Secret,
		now+this.appctx.Config.Jwt.Exp,
		now+this.appctx.Config.Jwt.Exp/2,
		po.ID,
		logincode,
		po.Role,
		po.IsSuperAdmin,
	); err != nil {
		return nil, this.ErrMsg("登陆失败", "jwt:"+err.Error())
	} else {

		refleshToken, err := this.newRefreshToken(po.ID)
		if err != nil {
			return nil, err
		}
		return mymodel.LogRep{
			token,
			now + this.appctx.Config.Jwt.Exp/2,
			refleshToken,
		}, nil

	}
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

func (this *Login) newLoginCode(id int64) (string, error) {
	//登陆处理
	//登陆code 控制唯一登陆有效及踢人
	var logincode string
	if logincode = uuid.New().String(); logincode == "" {
		return "", this.ErrMsg("logincode is empty", "登陆失败")
	} else {
		logincode = strings.Split(logincode, "-")[0]
		if r := this.appctx.Redis.Set(context.Background(), mymodel.GetAdminRedisLoginId(this.appctx.Config.App.Name, int(id)), logincode, 0); r.Err() != nil {
			return "", this.ErrMsg("redis:"+r.Err().Error(), "登陆失败")
		}
	}
	return logincode, nil
}

// 刷新token生成
func (this *Login) newRefreshToken(id int64) (string, error) {
	var refreshToken string
	if refreshToken = uuid.New().String(); refreshToken == "" {
		return "", this.ErrMsg("登陆失败", "refreshToken is empty")
	} else {
		if r := this.appctx.Redis.Set(context.Background(), mymodel.GetAdminRedisRefreshTokenId(this.appctx.Config.App.Name, int(id)), refreshToken, time.Second*time.Duration(this.appctx.Config.Jwt.Exp)); r.Err() != nil {
			return "", this.ErrMsg("redis:"+r.Err().Error(), "登陆失败")
		} else {
			return refreshToken, nil
		}
	}
}
