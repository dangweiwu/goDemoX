package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/auth/authmodel"
	"DEMOX_ADMINAUTH/internal/app/role/rolemodel"
	"DEMOX_ADMINAUTH/internal/pkg/dbtype"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAetAuth(t *testing.T) {
	app := newApp()
	defer app.Close()
	authpo := &authmodel.AuthPo{
		Code:   "auth1",
		Name:   "name1",
		Api:    "api1",
		Method: "method1",
		Kind:   "0",
	}
	app.Db.Create(authpo)

	role := NewRole()
	role.Auth = dbtype.List[string]{"auth1"}
	app.Db.Create(role)

	form := &rolemodel.RoleAuthForm{Auth: dbtype.List[string]{authpo.Code}}

	ser := app.Put(fmt.Sprintf("/api/role/auth/%d", role.ID), form).Do()
	assert.Equal(t, 200, ser.GetCode())

	//ser := testtool.NewTestServer(SerCtx, "PUT", fmt.Sprintf("/api/role/auth/%d", po1.ID), bytes.NewBuffer(body)).SetAuth(my.AccessToken).Do()
	//if assert.Equal(t, 200, ser.GetCode(), "%d:%s", ser.GetCode(), ser.GetBody()) {
	//	po := &rolemodel.RolePo{}
	//	err := SerCtx.Db.Model(po).Where("id=?", po1.ID).Take(po).Error
	//	assert.Nil(t, err)
	//	bts1, _ := json.Marshal(form.Auth)
	//	bts2, _ := json.Marshal(po.Auth)
	//	assert.Equal(t, bts1, bts2)
	//}
}
