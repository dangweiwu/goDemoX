package api

import (
	"github.com/gin-gonic/gin"
	"goDemoX/internal/app/role/rolemodel"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg/api/hd"
	"goDemoX/internal/pkg/api/query"
	"goDemoX/internal/router"
	"gorm.io/gorm"
)

type RoleQuery struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewRoleQuery(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &RoleQuery{hd.NewHd(c), c, appctx}
}

// Do
// @api 	| role | 4 | 角色查询
// @path 	| /api/role
// @method 	| GET
// @header 	|n Authorization |d token |e tokenstring |c 鉴权 |t string
// @query   |n limit   |d 条数 |e 10 |t int
// @query   |n current |d 页码 |e 1  |t int
// @query 	|n code |d 角色编码 |t string
// @query   |n name |d 角色名称
// @response | query.QueryResult | 200 Response
// @response | query.Page | Page定义
// @response | rolemodel.RolePo | []Data 定义
func (this *RoleQuery) Do() error {

	data, err := this.Query()
	if err != nil {
		return err
	} else {
		this.Rep(data)
		return nil
	}
}

var QueryRule = map[string]string{
	"code": "like",
	"name": "like",
}

func (this *RoleQuery) Query() (interface{}, error) {
	po := &rolemodel.RolePo{}
	pos := []rolemodel.RolePo{}
	q := query.NewQuery(this.ctx, this.appctx.Db, QueryRule, po, &pos)
	q.SetOrder(func(db *gorm.DB) (r *gorm.DB) {
		db = db.Order("order_num")
		return db
	})
	return q.Do()
}
