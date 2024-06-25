package api_test

import (
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

func refreshEnv(t *testing.T) (*testapp.TestApp, *adminmodel.AdminPo) {
	app, err := testapp.NewTestApp()
	if err != nil {
		panic(err)
	}

	app.GetUid = func(ctx *gin.Context) (int64, error) {
		return jwtx.GetUid(ctx)
	}

	app.RegDb(&adminmodel.AdminPo{})
	user := NewUser()
	app.Db.Create(user)
	app.RegRoute(func(engine *gin.Engine) {
		my.Route(router.NewTestRouter(engine, app.AppContext), app.AppContext)
	})

	return app, user

}

func TestRefreshToken(t *testing.T) {
	app, user := refreshEnv(t)
	defer app.Close()
	logform := &mymodel.LoginForm{user.Account, password}
	ser := app.Post("/api/login", logform).Do()
	if !assert.Equal(t, 200, app.GetCode(), "login:%d:%s", app.GetCode(), app.GetBody()) {
		return
	}
	//fmt.Println(ser.GetBody())
	rep := &mymodel.LogRep{}
	ser.ResponseObj(rep)
	assert.NotEmpty(t, rep.RefreshToken, "refreshtoken")
	assert.NotEmpty(t, rep.AccessToken, "accesstoken")
	assert.NotEmpty(t, rep.TokenExp, "TokenExp")

	ser = app.Post("/api/token/refresh", nil).SetToken(rep.RefreshToken).Do()
	//fmt.Println(ser.GetBody())
	if !assert.Equal(t, 200, app.GetCode(), "refresh:%d:%s", app.GetCode(), app.GetBody()) {
		return
	}
	ser.ResponseObj(rep)
	assert.NotEmpty(t, rep.RefreshToken, "refreshtoken")
	assert.NotEmpty(t, rep.AccessToken, "accesstoken")
	assert.NotEmpty(t, rep.TokenExp, "TokenExp")

}
