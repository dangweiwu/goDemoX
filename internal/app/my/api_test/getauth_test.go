package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/app/my"
	"DEMOX_ADMINAUTH/internal/app/role/rolemodel"
	"DEMOX_ADMINAUTH/internal/ctx/testapp"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/pkg/dbtype"
	"DEMOX_ADMINAUTH/internal/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetAuthEnv(t *testing.T) (*testapp.TestApp, *adminmodel.AdminPo) {
	user := NewUser()
	app, err := testapp.NewTestApp()
	if err != nil {
		panic(err)
	}
	app.RegDb(&adminmodel.AdminPo{}, &rolemodel.RolePo{})
	app.RegRoute(func(engine *gin.Engine) {
		my.Route(router.NewTestBaseRouter(engine, app.AppContext), app.AppContext)
	})

	//添加auth
	rolePo := &rolemodel.RolePo{}
	rolePo.Auth = dbtype.List[string]{"auth1", "auth2"}
	rolePo.Code = "admin"
	if r := app.Db.Create(rolePo); r.Error != nil {
		panic(r.Error)
	}

	user.Role = "admin"
	if r := app.Db.Create(user); r.Error != nil {
		panic(r.Error)
	}
	app.GetUid = func(ctx *gin.Context) (int64, error) {
		return user.ID, nil
	}

	app.GetRole = func(ctx *gin.Context) (string, error) {
		return user.Role, nil
	}

	return app, user
}

func TestGetAuth(t *testing.T) {
	fmt.Println("=============auth start")
	defer fmt.Println("end===============auth")
	app, _ := GetAuthEnv(t)
	defer app.Close()
	ser := app.Get("/api/my-auth").Do()
	if !assert.Equal(t, 200, ser.GetCode()) {
		fmt.Println(ser.GetBody())
		return
	}

	auths := hd.Response{Data: []string{}}
	err := ser.ResponseObj(&auths)
	if !assert.Nil(t, err) {
		fmt.Println(ser.GetBody())
		return
	}
	if !assert.NotNil(t, auths.Data) {
		fmt.Println(auths)
		return
	}

	assert.Equal(t, 2, len(auths.Data.([]interface{})))

}
