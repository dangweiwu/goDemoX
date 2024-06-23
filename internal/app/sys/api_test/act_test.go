package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/sys/sysmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAct(t *testing.T) {

	app := newApp()
	defer app.Close()

	form := sysmodel.SysActForm{}
	form.Name = "trace"
	form.Act = "1"

	ser := app.Put("/api/sys", form).Do()
	assert.Equal(t, 200, ser.GetCode(), ser.GetBody())

	form = sysmodel.SysActForm{}
	form.Name = "metric"
	form.Act = "1"

	ser = app.Put("/api/sys", form).Do()
	assert.Equal(t, 200, ser.GetCode(), ser.GetBody())

	form = sysmodel.SysActForm{}
	form.Name = "trace"
	form.Act = "0"

	ser = app.Put("/api/sys", form).Do()
	assert.Equal(t, 200, ser.GetCode(), ser.GetBody())

	form = sysmodel.SysActForm{}
	form.Name = "metric"
	form.Act = "0"

	ser = app.Put("/api/sys", form).Do()
	assert.Equal(t, 200, ser.GetCode(), ser.GetBody())

}
