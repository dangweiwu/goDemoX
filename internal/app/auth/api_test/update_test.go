package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/auth/authmodel"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthUpdate(t *testing.T) {

	app := newApp()
	app.Close()
	po := NewAuth()
	app.Db.Create(po)
	form := authmodel.AuthUpdateForm{Name: "创建1", OrderNum: 1002, Api: "/api/auth1", Method: "PUT", Kind: "0"}

	ser := app.Put(fmt.Sprintf("/api/auth/%d", po.ID), form).Do()
	if !assert.Equal(t, 200, ser.GetCode(), ser.GetBody()) {
		return
	}

	uppo := &authmodel.AuthPo{}
	app.Db.Take(uppo)

	assert.Equal(t, po.ID, uppo.ID)
	assert.Equal(t, form.Name, uppo.Name)
	assert.Equal(t, form.Api, uppo.Api)
	assert.Equal(t, form.Method, uppo.Method)
	assert.Equal(t, form.Kind, uppo.Kind)
	assert.Equal(t, form.OrderNum, uppo.OrderNum)

}
