package adminmodel

import (
	"errors"
	"fmt"
	"goDemoX/internal/pkg/jwtx"
	"time"
)

// access token时间小于refresh token时间
// 前端在access token时间到期之前进行刷新

type LoginOpt struct {
	LoginCode  string
	Secret     string
	Exp        int64
	RefreshExp int64
}

func (this *LoginOpt) Valid() error {
	if len(this.Secret) == 0 {
		return errors.New("jwt need secret")
	}
	if this.Exp == 0 {
		this.Exp = 24 * 60 * 60
	}

	if this.RefreshExp == 0 || this.RefreshExp < this.Exp {
		this.RefreshExp = this.Exp * 3 //3
	}

	return nil
}

// 登录逻辑
func (this AdminPo) Login(opt *LoginOpt) (token, refreshToken string, err error) {
	fmt.Println("============++++++++")
	if err = opt.Valid(); err != nil {
		return
	}

	now := time.Now().Unix()

	token, err = jwtx.GenToken(
		opt.Secret,
		now+opt.Exp, //过期时间
		this.ID,
		opt.LoginCode,
		this.Role,
		this.IsSuperAdmin,
	)

	if err != nil {
		err = fmt.Errorf("GenTokenErr %v", err)
		return
	}

	refreshToken, err = jwtx.GenToken(
		opt.Secret,
		now+opt.RefreshExp, //刷新时间
		this.ID,
		opt.LoginCode,
		this.Role,
		this.IsSuperAdmin,
	)
	if err != nil {
		err = fmt.Errorf("GenRefreshTokenErr %v", err)
		return
	}
	return
}
