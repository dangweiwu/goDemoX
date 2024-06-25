package api_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"goDemoX/internal/app/auth/authmodel"
	"testing"
)

func TestAuthCreate(t *testing.T) {

	app := newApp()
	defer app.Close()

	auth := NewAuth()
	ser := app.Post("/api/auth", auth).Do()
	if !assert.Equal(t, ser.GetCode(), 200, "%d:%s", ser.GetCode(), ser.GetBody()) {
		fmt.Println(t, 200, ser.GetCode())
		return

	}
	po := &authmodel.AuthPo{}
	app.Db.Where("code=?", auth.Code).Take(po)
	assert.Equal(t, po.Name, auth.Name, "name not equal")
	assert.Equal(t, po.Method, auth.Method, "method not equal")
	assert.Equal(t, po.Api, auth.Api, "api not equal")
	assert.Equal(t, po.OrderNum, auth.OrderNum, "order not equal")
	assert.Equal(t, po.Kind, auth.Kind, "kind not equal")

}
