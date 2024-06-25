package api_test

import (
	"goDemoX/internal/app/admin/adminmodel"
	"goDemoX/internal/pkg"
	"goDemoX/internal/pkg/dbtype"
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
