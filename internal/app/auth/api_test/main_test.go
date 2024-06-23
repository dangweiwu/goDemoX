package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/auth"
	"DEMOX_ADMINAUTH/internal/app/auth/authmodel"
	"DEMOX_ADMINAUTH/internal/app/role/rolemodel"
	"DEMOX_ADMINAUTH/internal/ctx/testapp"
	"DEMOX_ADMINAUTH/internal/pkg/dbtype"
	"DEMOX_ADMINAUTH/internal/router"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestMain(m *testing.M) {

	m.Run()
}

func NewRole() *rolemodel.RolePo {
	return &rolemodel.RolePo{
		Base:     dbtype.Base{ID: 1},
		Code:     "role1",
		Name:     "角色1",
		OrderNum: 1,
		Status:   "1",
		Memo:     "这是memo",
		Auth:     dbtype.List[string]{"1", "2"},
	}
}

func NewAuth() *authmodel.AuthPo {
	return &authmodel.AuthPo{
		Base:     dbtype.Base{ID: 1},
		Code:     "auth1",
		Name:     "权限1",
		OrderNum: 1,
		Kind:     "0",
		ParentId: 0,
		Api:      "/api/v1/auth/auth",
	}
}

func newApp() *testapp.TestApp {
	app, err := testapp.NewTestApp()
	if err != nil {
		panic(err)
	}
	app.RegDb(&rolemodel.RolePo{}, &authmodel.AuthPo{})
	app.RegRoute(func(engine *gin.Engine) {
		auth.Route(router.NewTestBaseRouter(engine, app.AppContext), app.AppContext)
	})
	if err := app.InitAuthCheckout(); err != nil {
		panic(err)
	}
	return app
}
