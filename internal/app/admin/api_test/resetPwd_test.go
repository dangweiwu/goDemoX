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

func resetPwdEnv() *testapp.TestApp {
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

func TestResetPwd(t *testing.T) {
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
	fmt.Println("user.id", user.ID)
	ser := app.Put(fmt.Sprintf("/api/admin/resetpwd/%d", user.ID), nil).Do()
	if assert.Equal(t, 200, ser.GetCode()) {
		assert.Len(t, ser.GetBody(), 17, "resetpwd-len:17")
		fmt.Println(ser.GetBody())
		return
	}

}
