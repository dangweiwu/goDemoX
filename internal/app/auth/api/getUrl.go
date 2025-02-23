package api

import (
	"github.com/gin-gonic/gin"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg/api/hd"
	"goDemoX/internal/pkg/fullurl"
	"goDemoX/internal/router"
)

/*
获取全部url
*/
type GetFullUrl struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewGetFullUrl(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &GetFullUrl{hd.NewHd(c), c, appctx}
}

// Do
// @api 	| auth | 5 | 获取所有URL | 创建修改权限时url参数从这获取
// @path 	| /api/auth
// @method 	| GET
// @header 	|n Authorization |d token |e tokenstring |c 鉴权 |t string
// @tbtitle | 200 Response
// @tbrow |n data |d 权限列表 |t []string |c 列表数据 |e ['/api/admin']
func (this *GetFullUrl) Do() error {
	this.Hd.Rep(hd.Response{fullurl.NewFullUrl().GetUrl()})
	return nil
}
