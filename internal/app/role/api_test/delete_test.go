package api_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoleDelete(t *testing.T) {
	app := newApp()
	defer app.Close()
	role1 := NewRole()
	app.Db.Create(role1)
	role2 := NewRole()
	role2.Code = "role2"
	role2.ID = 2
	app.Db.Create(role2)

	ser := app.Delete(fmt.Sprintf("/api/role/%d", role1.ID)).Do()
	assert.Equal(t, 200, ser.GetCode(), "%d:%s", ser.GetCode(), ser.GetBody())

	ser = app.Delete(fmt.Sprintf("/api/role/%d", role2.ID)).Do()
	assert.Equal(t, 200, ser.GetCode(), "%d:%s", ser.GetCode(), ser.GetBody())

}
