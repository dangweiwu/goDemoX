package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"goDemoX/internal/app/my/mymodel"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg/api/hd"
	"goDemoX/internal/router"
)

/*
退出登陆
*/
type LogOut struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewLogOut(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &LogOut{hd.NewHd(c), c, appctx}
}

// Do
// @api     | me | 2 | 登出
// @path 	| /api/logout
// @method 	| POST
// @header  |n Authorization |d token |t string |c 鉴权
// @tbtitle | 200 Response
// @tbrow   |n data |e ok |c 成功 |t string
func (this *LogOut) Do() error {
	this.Logout()
	this.Hd.RepOk()
	return nil
}

func (this *LogOut) Logout() error {
	//获取id
	id, err := this.appctx.GetUid(this.ctx)
	if err != nil {
		return nil
	}
	//删除logincode
	this.appctx.Redis.Del(context.Background(), mymodel.GetAdminRedisLoginId(int(id)))
	return nil
}
