package api_test

import (
	"github.com/gin-gonic/gin"
	"goDemoX/internal/app/sys"
	"goDemoX/internal/ctx/testapp"
	"goDemoX/internal/router"
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
