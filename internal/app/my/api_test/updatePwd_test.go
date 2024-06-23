package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/app/my"
	"DEMOX_ADMINAUTH/internal/app/my/mymodel"
	"DEMOX_ADMINAUTH/internal/ctx/testapp"
	"DEMOX_ADMINAUTH/internal/pkg"
	"DEMOX_ADMINAUTH/internal/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func updatemypwdEnv() (*testapp.TestApp, *adminmodel.AdminPo) {
	app, err := testapp.NewTestApp()
	if err != nil {
		panic(err)
	}
	app.RegDb(&adminmodel.AdminPo{})
	user := NewUser()
	app.Db.Create(user)
	app.GetUid = func(ctx *gin.Context) (int64, error) {
		return user.ID, nil
	}
	app.RegRoute(func(engine *gin.Engine) {
		my.Route(router.NewTestBaseRouter(engine, app.AppContext), app.AppContext)
	})
	return app, user

}

func TestUpdateMyPwd(t *testing.T) {
	var (
		newpwd = "a123456"
	)
	app, user := updatemypwdEnv()
	defer app.Close()

	form := &mymodel.PasswordForm{password, newpwd}

	ser := app.Put("/api/my/password", form).Do()
	if !assert.Equal(t, 200, ser.GetCode()) {
		fmt.Println(ser.GetBody())
		return
	}

	newuser := &adminmodel.AdminPo{}
	app.Db.Model(newuser).Where("id=?", user.ID).Take(newuser)
	assert.Equal(t, pkg.GetPassword(newpwd), newuser.Password)

}
