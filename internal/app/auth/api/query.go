package api

import (
	"github.com/gin-gonic/gin"
	"goDemoX/internal/app/auth/authmodel"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg/api/hd"
	"goDemoX/internal/router"
)

type AuthQuery struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewAuthQuery(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &AuthQuery{hd.NewHd(c), c, appctx}
}

// Do
// @api 	| auth | 3 | 权限查询
// @path 	| /api/auth
// @method 	| GET
// @header 	|n Authorization |d token |e tokenstring |c 鉴权 |t string
// @query 	|n kind |d 类型 |e 0 |t string |c 0:按钮 1:页面
// @response | hd.Response | 200 Response
// @response | authmodel.AuthVo | Data定义
func (this *AuthQuery) Do() error {

	data, err := this.Query()
	if err != nil {
		return err
	} else {
		this.Rep(hd.Response{data})
		return nil
	}
}

var QueryRule = map[string]string{
	"kind": "like",
}

func (this *AuthQuery) Query() (interface{}, error) {
	po := &authmodel.AuthVo{}
	pos := []authmodel.AuthVo{}

	kind := this.ctx.Query("kind")
	if kind != "" {
		if err := this.appctx.Db.Model(po).Where("parent_id=0 and kind= ?", kind).Preload("Children", "kind=?", kind).Order("order_num").Find(&pos).Error; err != nil {
			return nil, err
		}
	} else {
		if err := this.appctx.Db.Model(po).Where("parent_id=0").Preload("Children").Order("order_num").Find(&pos).Error; err != nil {
			return nil, err
		}
	}

	return pos, nil
}
