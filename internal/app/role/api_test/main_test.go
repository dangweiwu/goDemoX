package api_test

import (
	"github.com/gin-gonic/gin"
	"goDemoX/internal/app/auth/authmodel"
	"goDemoX/internal/app/role"
	"goDemoX/internal/app/role/rolemodel"
	"goDemoX/internal/ctx/testapp"
	"goDemoX/internal/pkg/dbtype"
	"goDemoX/internal/router"
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

func newApp() *testapp.TestApp {
	app, err := testapp.NewTestApp()
	if err != nil {
		panic(err)
	}
	app.RegDb(&rolemodel.RolePo{}, &authmodel.AuthPo{})
	app.RegRoute(func(engine *gin.Engine) {
		role.Route(router.NewTestBaseRouter(engine, app.AppContext), app.AppContext)
	})
	if err := app.InitAuthCheckout(); err != nil {
		panic(err)
	}
	return app
}
