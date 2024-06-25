package api_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"goDemoX/internal/app/admin/adminmodel"
	"goDemoX/internal/app/my"
	"goDemoX/internal/ctx/testapp"
	"goDemoX/internal/router"
	"log"
	"testing"
)

func infoEnv() (*testapp.TestApp, *adminmodel.AdminPo) {
	app, err := testapp.NewTestApp()
	if err != nil {
		panic(err)
	}
	app.RegDb(&adminmodel.AdminPo{})
	app.GetUid = func(ctx *gin.Context) (int64, error) {
		return 1, nil
	}
	app.RegRoute(func(engine *gin.Engine) {
		my.Route(router.NewTestBaseRouter(engine, app.AppContext), app.AppContext)
	})
	user := NewUser()
	app.Db.Create(user)
	return app, user

}

func TestInfo(t *testing.T) {
	fmt.Println("=============info start")
	defer fmt.Println("end===============info")
	app, user := infoEnv()
	defer app.Close()
	ser := app.Get("/api/my").Do()
	if assert.Equal(t, 200, ser.GetCode()) {
		userPo := &adminmodel.AdminPo{}
		log.Println(ser.GetBody())
		err := ser.ResponseObj(userPo)
		assert.Nil(t, err)
		assert.Equal(t, user.Name, userPo.Name)
		assert.Equal(t, user.Account, userPo.Account)
		assert.Equal(t, user.Phone, userPo.Phone)
		assert.Equal(t, user.Email, userPo.Email)
		assert.Equal(t, user.Role, userPo.Role)
		assert.Equal(t, user.Memo, userPo.Memo)
		assert.Equal(t, "", userPo.Status)
		assert.Equal(t, user.IsSuperAdmin, userPo.IsSuperAdmin)
		assert.Equal(t, user.ID, userPo.ID)
	}

}
