package api

import (
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/pkg/api/query"
	"DEMOX_ADMINAUTH/internal/router"
	"github.com/gin-gonic/gin"
)

type AdminQuery struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewAdminQuery(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &AdminQuery{hd.NewHd(c), c, appctx}
}

// Do
// @api 	| admin | 4 | 查询用户
// @path 	| /api/admin
// @method 	| GET
// @header 	|n Authorization |d token |e tokenstring |c 鉴权 |t string
// @query   |n limit   |d 条数 |e 10 |t int
// @query   |n current |d 页码 |e 1  |t int
// @query 	|n account |d 账号 |e admin | t string
// @query   |n phone   |d 手机号 |e 12345678911 |t int
// @query   |n email   |d email
// @query   |n name    |d 姓名
// @response | query.QueryResult | 200 Response
// @response | query.Page | Page定义
// @response | adminmodel.AdminVo | []Data 定义
func (this *AdminQuery) Do() error {
	data, err := this.Query()
	if err != nil {
		return err
	} else {
		this.Rep(data)
		return nil
	}
}

var QueryRule = map[string]string{
	"account": "like",
	"phone":   "like",
	"email":   "like",
	"name":    "like",
}

func (this *AdminQuery) Query() (interface{}, error) {

	po := &adminmodel.AdminVo{}
	pos := []adminmodel.AdminVo{}
	q := query.NewQuery(this.ctx, this.appctx.Db, QueryRule, po, &pos)
	/*
		q.SetWhere(func(db *gorm.DB) (r *gorm.DB) {
			db = db.Preload("RolePo")
			return q.Where(db)
		})
	*/
	return q.Do()
}
