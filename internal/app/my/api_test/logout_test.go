package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/app/my"
	"DEMOX_ADMINAUTH/internal/app/my/mymodel"
	"DEMOX_ADMINAUTH/internal/app/my/myserver"
	"DEMOX_ADMINAUTH/internal/ctx/testapp"
	"DEMOX_ADMINAUTH/internal/router"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

var logcode string

func logoutEnv() (*testapp.TestApp, *adminmodel.AdminPo) {
	app, err := testapp.NewTestApp()
	if err != nil {
		panic(err)
	}
	app.RegDb(&adminmodel.AdminPo{})
	user := NewUser()
	_, err = myserver.NewLogCode(user.ID, app.AppContext.Redis)
	if err != nil {
		panic(err)
	}
	//logcode = logcode
	app.Db.Create(user)
	app.GetUid = func(ctx *gin.Context) (int64, error) {
		return user.ID, nil
	}
	app.RegRoute(func(engine *gin.Engine) {
		my.Route(router.NewTestBaseRouter(engine, app.AppContext), app.AppContext)
	})
	return app, user
}

func TestLogout(t *testing.T) {
	app, user := logoutEnv()
	defer app.Close()
	r := app.Redis.Get(context.Background(), mymodel.GetAdminRedisLoginId(int(user.ID)))
	a, b := r.Result()
	assert.NotEmpty(t, a)
	assert.Nil(t, b)
	ser := app.Post("/api/logout", nil).Do()
	if assert.Equal(t, 200, ser.GetCode()) {
		r := app.Redis.Get(context.Background(), mymodel.GetAdminRedisLoginId(int(user.ID)))
		a, b := r.Result()
		assert.Empty(t, a)
		assert.Equal(t, redis.Nil, b)
	}
}
