package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/sys"
	"DEMOX_ADMINAUTH/internal/ctx/testapp"
	"DEMOX_ADMINAUTH/internal/router"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func newApp() *testapp.TestApp {
	app, err := testapp.NewTestApp()
	if err != nil {
		panic(err)
	}
	app.RegRoute(func(engine *gin.Engine) {
		sys.Route(router.NewTestBaseRouter(engine, app.AppContext), app.AppContext)
	})
	app.Config.Metric.Enable = true
	app.Config.Trace.Enable = true
	app.InitTrace()
	app.InitMetric()
	return app
}
