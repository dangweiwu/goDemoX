package api_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInfo(t *testing.T) {
	app := newApp()
	defer app.Close()

	ser := app.Get("/api/sys").Do()
	assert.Equal(t, 200, ser.GetCode(), ser.GetBody())
}
