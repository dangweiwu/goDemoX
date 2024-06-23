package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/admin"
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/ctx/testapp"
	"DEMOX_ADMINAUTH/internal/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func updateEnv() *testapp.TestApp {
	app, err := testapp.NewTestApp()
	if err != nil {
		panic(err)
	}
	app.RegDb(&adminmodel.AdminPo{})
	app.RegRoute(func(engine *gin.Engine) {
		admin.Route(router.NewTestBaseRouter(engine, app.AppContext), app.AppContext)
	})
	return app
}
func TestAdminUpdate(t *testing.T) {
	app := resetPwdEnv()
	defer app.Close()
	my := NewUser()
	app.Db.Create(my)
	app.GetUid = func(ctx *gin.Context) (int64, error) {
		fmt.Println("my.id", my.ID)
		return my.ID, nil
	}

	user := NewUser()
	user.ID = 2
	user.Account = "admin2"
	app.Db.Create(user)

	upform := &adminmodel.AdminUpdateForm{Phone: "12312312311", Name: "namechang", Status: "0",
		Memo: "memochange", Email: "email@qq.com", IsSuperAdmin: "1"}

	ser := app.Put(fmt.Sprintf("/api/admin/%d", user.ID), upform).Do()
	if !assert.Equal(t, 200, ser.GetCode()) {
		return
	}
	newPo := &adminmodel.AdminPo{}
	app.Db.Where("account=?", user.Account).Take(newPo)
	assert.Equal(t, upform.Phone, newPo.Phone, "update:phone")
	assert.Equal(t, upform.Name, newPo.Name, "update:name")
	assert.Equal(t, upform.Status, newPo.Status, "update:status")
	assert.Equal(t, upform.Memo, newPo.Memo, "update:memo")
	assert.Equal(t, upform.Email, newPo.Email, "update:email")
	assert.Equal(t, upform.IsSuperAdmin, newPo.IsSuperAdmin, "update:update")

}
