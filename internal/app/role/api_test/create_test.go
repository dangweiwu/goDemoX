package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/role/rolemodel"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRole(t *testing.T) {

	app := newApp()
	defer app.Close()
	form := &rolemodel.RoleForm{Code: "admin", Name: "系统管理员", OrderNum: 1, Status: "1", Memo: "this is memo"}
	ser := app.Post("/api/role", form).Do()

	if !assert.Equal(t, 200, ser.GetCode()) {
		fmt.Println(ser.GetBody())
	}
	po := &rolemodel.RolePo{}
	app.Db.Where("code=?", form.Code).Take(po)
	assert.Equal(t, po.Name, form.Name)
	assert.Equal(t, po.Status, form.Status)
	assert.Equal(t, po.OrderNum, form.OrderNum)
	assert.Equal(t, po.Code, form.Code)
	assert.Equal(t, po.Memo, form.Memo)

	ser = app.Post("/api/role", form).Do()

	if assert.Equal(t, 400, ser.GetCode(), "%d:%s", ser.GetCode(), ser.GetBody()) {
		assert.Contains(t, ser.GetBody(), "角色编码已存在")
	}

}
