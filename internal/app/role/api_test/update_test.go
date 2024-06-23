package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/role/rolemodel"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateRole(t *testing.T) {
	app := newApp()
	defer app.Close()
	role := NewRole()
	app.Db.Create(role)

	form := &rolemodel.RoleForm{Name: "nameupdate", OrderNum: 2, Status: "0", Memo: "updateMemo"}

	ser := app.Put(fmt.Sprintf("/api/role/%d", role.ID), form).Do()
	if !assert.Equal(t, 200, ser.GetCode()) {
		fmt.Println(ser.GetBody())
		return
	}
	aftpo := &rolemodel.RolePo{}
	app.Db.Where("id=?", role.ID).Take(aftpo)
	assert.Equal(t, form.Name, aftpo.Name)
	assert.Equal(t, form.OrderNum, aftpo.OrderNum)
	assert.Equal(t, form.Status, aftpo.Status)
	assert.Equal(t, form.Memo, aftpo.Memo)

}
