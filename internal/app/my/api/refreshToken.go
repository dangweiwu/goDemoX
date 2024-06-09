package api

/*
token刷新
*/
import (
	"DEMOX_ADMINAUTH/internal/app/my/mymodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/pkg/jwtx"
	"DEMOX_ADMINAUTH/internal/router"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type RefreshToken struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewRefreshToken(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &RefreshToken{hd.NewHd(c), c, appctx}
}

// Do
// @api     | me | 6 | 刷新token
// @path 	| /api/token/refresh
// @method 	| POST
// @header   |n Authorization |d token  |t string |c 鉴权
// @form     | mymodel.RefreshTokeForm
// @response | mymodel.LogRep |200 Response
// @response | hd.ErrResponse | 401 Response
// @tbtitle | Msg 数据
// @tbrow   |n msg |d refreshtoken已失效
func (this *RefreshToken) Do() error {

	form := &mymodel.RefreshTokeForm{}
	if err := this.Bind(form); err != nil {
		return err
	}

	uid, err := jwtx.GetUid(this.ctx)
	if err != nil {
		return err
	}

	//刷新token检验
	r, err := this.appctx.Redis.Get(context.Background(), mymodel.GetAdminRedisRefreshTokenId(this.appctx.Config.App.Name, int(uid))).Result()
	if err != nil {
		return err
	}

	if r != form.RefreshToken {
		this.ctx.JSON(401, hd.ErrMsg("refreshtoken已失效", ""))
		return nil
	}
	//数据源
	data, err := this.RefreshToken(uid)
	if err != nil {
		return err
	}

	this.Rep(data)
	return nil
}

func (this *RefreshToken) RefreshToken(uid int64) (interface{}, error) {
	// logincode, err := this.newCode(uid)

	logincode, err := jwtx.GetCode(this.ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	role, _ := jwtx.GetRole(this.ctx)
	issuper, _ := jwtx.GetIsSuper(this.ctx)
	_issuper := ""
	if issuper {
		_issuper = "1"
	} else {
		_issuper = "0"
	}
	if token, err := jwtx.GenToken(
		this.appctx.Config.Jwt.Secret,
		now+this.appctx.Config.Jwt.Exp,
		now+this.appctx.Config.Jwt.Exp/2,
		uid,
		logincode,
		role,
		_issuper,
	); err != nil {
		return nil, this.ErrMsg("刷新失败", "jwt:"+err.Error())
	} else {
		// this.ctx.Header("Authorization", token)
		newRefreshToken, err := this.newRefreshToken(uid)
		if err != nil {
			return "", err
		}
		return mymodel.LogRep{
			token,
			now + this.appctx.Config.Jwt.Exp/2,
			newRefreshToken,
		}, nil
	}
}

// 刷新token生成
func (this *RefreshToken) newRefreshToken(id int64) (string, error) {
	var refreshToken string
	if refreshToken = uuid.New().String(); refreshToken == "" {
		return "", this.ErrMsg("刷新失败", "refreshToken is empty")
	} else {
		if r := this.appctx.Redis.Set(context.Background(), mymodel.GetAdminRedisRefreshTokenId(this.appctx.Config.App.Name, int(id)), refreshToken, time.Second*time.Duration(this.appctx.Config.Jwt.Exp)); r.Err() != nil {
			return "", this.ErrMsg("刷新失败", "redis:"+r.Err().Error())
		} else {
			return refreshToken, nil
		}
	}
}
