package api_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoleQuery(t *testing.T) {
	app := newApp()
	defer app.Close()
	role1 := NewRole()
	app.Db.Create(role1)
	role2 := NewRole()
	role2.Name = "角色2"
	role2.Code = "role2"
	role2.ID = 2
	app.Db.Create(role2)

	ser := app.Get("/api/role").Do()
	if !assert.Equal(t, 200, ser.GetCode()) {
		fmt.Println(ser.GetBody())
		return
	}
	assert.Contains(t, ser.GetBody(), `"total":2`)

	ser = app.Get("/api/role?name=" + role1.Name).Do()

	if !assert.Equal(t, 200, ser.GetCode()) {
		fmt.Println(ser.GetBody())
		return
	}
	assert.Contains(t, ser.GetBody(), `"total":1`)

	ser = app.Get("/api/role?code=" + role1.Code).Do()
	if assert.Equal(t, 200, ser.GetCode(), "%d:%s", ser.GetCode(), ser.GetBody()) {
		assert.Contains(t, ser.GetBody(), `"total":1`)
	}
}
