package hd

import "github.com/gin-gonic/gin"

// @doc | hd.Response
type Response struct {
	Data interface{} `json:"data" doc:"|t any |c 响应数据 参考Data定义或说明"`
}

func Rep(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}

func RepOk(c *gin.Context) {
	c.JSON(200, Response{"ok"})
}
