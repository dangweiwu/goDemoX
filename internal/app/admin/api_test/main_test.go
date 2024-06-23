package api_test

import (
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/pkg"
	"DEMOX_ADMINAUTH/internal/pkg/dbtype"
	"testing"
)

func TestMain(m *testing.M) {

	m.Run()
}

var password = "123456"

func NewUser() *adminmodel.AdminPo {
	return &adminmodel.AdminPo{
		Base:         dbtype.Base{ID: 1},
		Role:         "abc",
		IsSuperAdmin: "1",
		Password:     pkg.GetPassword(password),
		Email:        "abc@qq.com",
		Status:       "1",
		Name:         "姓名",
		Account:      "admin",
		Phone:        "12222222222",
		Memo:         "这是memo",
	}
}
