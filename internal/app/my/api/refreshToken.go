package api

/*
token刷新
*/
import (
	"github.com/gin-gonic/gin"
	"goDemoX/internal/app/admin/adminmodel"
	"goDemoX/internal/app/my/mymodel"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg/api/hd"
	"goDemoX/internal/pkg/jwtx"
	"goDemoX/internal/router"
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

	//form := &mymodel.RefreshTokeForm{}
	//if err := this.Bind(form); err != nil {
	//	return err
	//}

	uid, err := this.appctx.GetUid(this.ctx)
	if err != nil {
		return err
	}

	po := &adminmodel.AdminPo{}
	if err := this.appctx.Db.Where("id=?", uid).First(po).Error; err != nil {
		return err
	}

	logcode, err := jwtx.GetCode(this.ctx)
	if err != nil {
		return err
	}

	logopt := adminmodel.LoginOpt{
		Exp:       this.appctx.Config.Jwt.Exp,
		LoginCode: logcode,
		Secret:    this.appctx.Config.Jwt.Secret,
	}
	token, refreshToken, err := po.Login(&logopt)
	if err != nil {
		return err
	}

	rep := &mymodel.LogRep{
		token,
		time.Now().Unix() + logopt.Exp - 60,
		refreshToken,
	}

	this.Rep(rep)
	return nil
}
