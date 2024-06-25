package api_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"goDemoX/internal/app/auth/authmodel"
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
