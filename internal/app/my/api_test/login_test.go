package api_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"goDemoX/internal/app/admin/adminmodel"
	"goDemoX/internal/app/my"
	"goDemoX/internal/app/my/mymodel"
	"goDemoX/internal/ctx/testapp"
	"goDemoX/internal/pkg/jwtx"
	"goDemoX/internal/router"
	"testing"
)

func LoginInitEnv() *testapp.TestApp {
	TestApp, err := testapp.NewTestApp()
	if err != nil {
		panic(err)
	}

	//初始化数据库
	TestApp.RegDb(&adminmodel.AdminPo{})
	//环境设置
	TestApp.GetUid = func(ctx *gin.Context) (int64, error) {
		return jwtx.GetUid(ctx)
	}
	//注册路由
	TestApp.RegRoute(func(engine *gin.Engine) {
		r := router.NewTestRouter(engine, TestApp.AppContext)
		my.Route(r, TestApp.AppContext)
	})

	return TestApp
}

func TestLogin(t *testing.T) {
	fmt.Println("=============login start")
	defer fmt.Println("end===============login")
	user := NewUser()
	testapp := LoginInitEnv()
	defer testapp.Close()
	testapp.Db.Create(user)

	form := &mymodel.LoginForm{user.Account, password}
	ser := testapp.Post("/api/login", form).Do()
	rep := &mymodel.LogRep{}
	ser.ResponseObj(rep)

	firstAccountToken := rep.AccessToken

	if assert.Equal(t, 200, ser.GetCode(), "login:%d:%s", ser.GetCode(), ser.GetBody()) {
		assert.NotEmpty(t, rep.AccessToken)
		assert.NotEmpty(t, rep.TokenExp)
		assert.NotEmpty(t, rep.RefreshToken)
	}

	//正常获取我的信息

	ser = testapp.Get("/api/my").SetToken(firstAccountToken).Do()
	assert.Equal(t, 200, ser.GetCode(), "my:info:%d:%s", ser.GetCode(), ser.GetBody())

	//二次登录
	ser = testapp.Post("/api/login", form).Do()

	//code失效校验
	ser = testapp.Get("/api/my").SetToken(firstAccountToken).Do()
	assert.Equal(t, 401, ser.GetCode(), "my:info:%d:%s", ser.GetCode(), ser.GetBody())

	//token 过期校验
	testapp.Config.Jwt.Exp = -1 //过期
	ser = testapp.Post("/api/login", form).Do()
	ser.ResponseObj(rep)
	token3 := rep.AccessToken

	ser = testapp.Get("/api/my").SetToken(token3).Do()
	ser.ResponseObj(rep)
	assert.Equal(t, 401, ser.GetCode(), "my:info:%d:%s", ser.GetCode(), ser.GetBody())

	//密码错误
	form = &mymodel.LoginForm{"admin", "123455"}
	testapp.Config.Jwt.Exp = 100
	ser = testapp.Post("/api/login", form).Do()
	if assert.Equal(t, 400, ser.GetCode(), "login:%d:%s", ser.GetCode(), ser.GetBody()) {
		assert.Contains(t, ser.GetBody(), "密码错误")
	}
}
