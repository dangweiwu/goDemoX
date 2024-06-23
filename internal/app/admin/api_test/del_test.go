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

func adminDelEnv() *testapp.TestApp {
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

func TestAdminDel(t *testing.T) {
	app := adminDelEnv()
	defer app.Close()
	user := NewUser()
	app.Db.Create(user)
	app.GetUid = func(ctx *gin.Context) (int64, error) {
		fmt.Println(user.ID)
		return user.ID, nil
	}
	user2 := NewUser()
	user2.Account = "admin2"
	user2.ID = 2
	if r := app.Db.Create(user2); r.Error != nil {
		panic(r.Error)
	}

	ser := app.Delete(fmt.Sprintf("/api/admin/%d", user2.ID)).Do()
	if !assert.Equal(t, 200, ser.GetCode()) {
		fmt.Println(ser.GetBody())
		return
	}
	users := []adminmodel.AdminPo{}
	app.Db.Find(&user)
	assert.Equal(t, 0, len(users), "del:user")

}
