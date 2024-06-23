package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/admin"
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/ctx/testapp"
	"DEMOX_ADMINAUTH/internal/pkg"
	"DEMOX_ADMINAUTH/internal/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func adminCreateEnv() *testapp.TestApp {
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

func TestAdminCreate(t *testing.T) {

	app := adminCreateEnv()
	defer app.Close()
	tests := []struct {
		name    string
		Method  string
		Target  string
		wantErr bool
		po      *adminmodel.AdminForm
	}{
		{"创建", "POST", "/api/admin", false,
			&adminmodel.AdminForm{Name: "dang", Phone: "12345678911", Account: "dang", Password: "123456", IsSuperAdmin: "0", Status: "1", Email: "abc1@qq.com"}},
		{"重复账号", "POST", "/api/admin", false,
			&adminmodel.AdminForm{Name: "dang", Phone: "12345678911", Account: "dang", Password: "123456", IsSuperAdmin: "0", Status: "1", Email: "abc2@qq.com"}},
		{"email格式", "POST", "/api/admin", false,
			&adminmodel.AdminForm{Name: "dang", Phone: "12345678912", Account: "dang3", Password: "123456", IsSuperAdmin: "0", Status: "1", Email: "abc"}},
	}
	for k, tt := range tests {
		ser := app.Post(tt.Target, tt.po).Do()

		switch k {
		case 0:
			if !assert.Equal(t, 200, ser.GetCode()) {
				fmt.Println(ser.GetBody())
				return
			}

			rpo := &adminmodel.AdminPo{}
			app.Db.Model(rpo).Take(rpo)
			fmt.Println(rpo)
			assert.Equal(t, rpo.Phone, tt.po.Phone)
			assert.Equal(t, rpo.Name, tt.po.Name)
			assert.Equal(t, rpo.Password, pkg.GetPassword(tt.po.Password))
			assert.Equal(t, rpo.IsSuperAdmin, tt.po.IsSuperAdmin)
			assert.Equal(t, rpo.Status, tt.po.Status)

		case 1:

			if assert.Equal(t, 400, ser.GetCode(), "%s:%s", tt.name, ser.GetBody()) {
				assert.Contains(t, ser.GetBody(), "账号已存在", "%s:%s", tt.name, "账号已存在")
			}

		case 2:
			if assert.Equal(t, 400, ser.GetCode(), "%s:%s", tt.name, ser.GetBody()) {
				assert.Contains(t, ser.GetBody(), "AdminForm.Email", "%s:%s", tt.name, "AdminForm.Email")
			}

		}

	}

}
