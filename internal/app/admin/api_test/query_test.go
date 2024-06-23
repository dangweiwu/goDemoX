package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/admin"
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/ctx/testapp"
	"DEMOX_ADMINAUTH/internal/router"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
查询测试
查询条件 account phone name email
*/

func queryEnv() *testapp.TestApp {
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

func TestAdminQuery(t *testing.T) {

	app := queryEnv()

	accountUser := &adminmodel.AdminPo{Account: "account1", Name: "name1", Phone: "12345678911", Email: "email1@qq.com"}
	phoneUser := &adminmodel.AdminPo{Account: "account2", Name: "name2", Phone: "12345678912", Email: "email2@qq.com"}
	nameUser := &adminmodel.AdminPo{Account: "account3", Name: "name3", Phone: "12345678913", Email: "email3@qq.com"}
	emailUser := &adminmodel.AdminPo{Account: "account4", Name: "name4", Phone: "12345678914", Email: "email4@qq.com"}
	app.Db.Create(accountUser)
	app.Db.Create(phoneUser)
	app.Db.Create(nameUser)
	app.Db.Create(emailUser)

	fmt.Println("=====================")
	pos := []adminmodel.AdminPo{}
	app.Db.Find(&pos)
	bts, _ := json.Marshal(pos)
	fmt.Println(string(bts))
	fmt.Println("=====================")

	//account test
	ser := app.Get("/api/admin?account=account1").Do()
	if assert.Equal(t, ser.GetCode(), 200, "%s:%s", "query:account=account1", ser.GetBody()) {
		assert.Contains(t, ser.GetBody(), "\"total\":1", "account1-total1")
		assert.Contains(t, ser.GetBody(), "account1", "account1-account1")
	}
	ser = app.Get("/api/admin?account=account").Do()
	if assert.Equal(t, ser.GetCode(), 200, "%s:%s", "query:account=account", ser.GetBody()) {
		assert.Contains(t, ser.GetBody(), "\"total\":4", "account-total4")
	}

	//name test
	ser = app.Get("/api/admin?name=name1").Do()
	if assert.Equal(t, ser.GetCode(), 200, "%s:%s", "query:name=mame1", ser.GetBody()) {
		assert.Contains(t, ser.GetBody(), "\"total\":1", "name1-total1")
		assert.Contains(t, ser.GetBody(), "name1", "name1-name1")
	}

	ser = app.Get("/api/admin?name=name").Do()
	if assert.Equal(t, ser.GetCode(), 200, "%s:%s", "query:name=name", ser.GetBody()) {
		assert.Contains(t, ser.GetBody(), "\"total\":4", "mame-total4")
	}
	ser = app.Get("/api/admin?phone=12345678911").Do()
	if assert.Equal(t, ser.GetCode(), 200, "%s:%s", "query:phone=1", ser.GetBody()) {
		assert.Contains(t, ser.GetBody(), "\"total\":1", "phone-total1")
		assert.Contains(t, ser.GetBody(), "12345678911", "phone-1")
	}

	ser = app.Get("/api/admin?email=email1@qq.com").Do()
	if assert.Equal(t, ser.GetCode(), 200, "%s:%s", "query:email=1", ser.GetBody()) {
		assert.Contains(t, ser.GetBody(), "\"total\":1", "email-total1")
		assert.Contains(t, ser.GetBody(), "email1@qq.com", "email-1")
	}

	ser = app.Get("/api/admin?account=12345").Do()
	if assert.Equal(t, ser.GetCode(), 200, "%s:%s", "query:noaccount:123456", ser.GetBody()) {
		assert.Contains(t, ser.GetBody(), "\"total\":0", "no-total0")
	}

	ser = app.Get("/api/admin").Do()
	if assert.Equal(t, ser.GetCode(), 200, "%s:%s", "query:all", ser.GetBody()) {
		assert.Contains(t, ser.GetBody(), "\"total\":4", "all-total:5")
	}
	ser = app.Get("/api/admin?account=acc&&name=nam&&phone=123&&email=ema").Do()
	if assert.Equal(t, ser.GetCode(), 200, "%s:%s", "query:all-union", ser.GetBody()) {
		assert.Contains(t, ser.GetBody(), "\"total\":4", "all-union-total:4")
	}

}
