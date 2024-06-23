package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/auth/authmodel"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthDel(t *testing.T) {
	app := newApp()
	defer app.Close()

	authpo := NewAuth()
	app.Db.Create(authpo)

	ser := app.Delete(fmt.Sprintf("/api/auth/%d", authpo.ID)).Do()
	if !assert.Equal(t, 200, ser.GetCode(), "%d:%s", ser.GetCode(), ser.GetBody()) {
		return
	}
	cnt := int64(0)
	app.Db.Model(&authmodel.AuthPo{}).Count(&cnt)
	assert.Equal(t, 0, int(cnt))

}
