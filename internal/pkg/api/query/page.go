package query

import "github.com/gin-gonic/gin"

const (
	LIMIT   = "limit"
	TOTAL   = "total"
	CURRENT = "current"
	//DB_ORDER     = "created_at desc"
)

var (
	LIMIT_CONF   = 10
	CURRENT_CONF = 1
)

// Page
// @doc | query.Page
type Page struct {
	Limit   int `json:"limit" form:"limit" doc:"|d 条数 |t int"`       // 每页条数
	Current int `json:"current" form:"current" doc:"|d 当前页码 |t int"` //当前页数
	Total   int `json:"total" doc:"|d 总数 |t int"`                    //总数
}

// 解析page 选项
func ParsePage(ctx *gin.Context) *Page {
	page := &Page{}
	ctx.ShouldBindQuery(page)

	if page.Limit == 0 {
		page.Limit = LIMIT_CONF
	}
	if page.Current == 0 {
		page.Current = CURRENT_CONF
	}
	return page
}
