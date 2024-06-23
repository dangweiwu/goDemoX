package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/app/my"
	"DEMOX_ADMINAUTH/internal/app/my/mymodel"
	"DEMOX_ADMINAUTH/internal/ctx/testapp"
	"DEMOX_ADMINAUTH/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func myupdateEnv() (*testapp.TestApp, *adminmodel.AdminPo) {
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

func TestMyUpdate(t *testing.T) {
	var (
		phone = "22222222222"
		name  = "name2"
		Memo  = "memo2"
		Email = "email2@qq.com"
	)
	app, user := myupdateEnv()
	defer app.Close()

	form := &mymodel.MyForm{Phone: phone, Name: name, Memo: Memo, Email: Email}
	ser := app.Put("/api/my", form).Do()
	if !assert.Equal(t, 200, ser.GetCode()) {
		return
	}
	_po := &adminmodel.AdminPo{}
	app.Db.Model(_po).Where("id=?", user.ID).Take(_po)

	if !assert.Equal(t, phone, _po.Phone) {
		return
	}

	if !assert.Equal(t, name, _po.Name) {
		return
	}
	if !assert.Equal(t, Memo, _po.Memo) {
		return
	}

	if !assert.Equal(t, Email, _po.Email) {
		return
	}

}
