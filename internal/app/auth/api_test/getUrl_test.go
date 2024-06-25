package api_test

import (
	"github.com/stretchr/testify/assert"
	"goDemoX/internal/pkg/fullurl"
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
