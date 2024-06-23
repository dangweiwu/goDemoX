package api_test

import (
	"DEMOX_ADMINAUTH/internal/pkg/fullurl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUrl(t *testing.T) {
	app := newApp()
	defer app.Close()

	fullurl.NewFullUrl().InitUrl(app.Engine)

	ser := app.Get("/api/allurl").Do()
	if !assert.Equal(t, 200, ser.GetCode(), ser.GetBody()) {
		return
	}

	assert.Contains(t, ser.GetBody(), "/api/auth")

}
