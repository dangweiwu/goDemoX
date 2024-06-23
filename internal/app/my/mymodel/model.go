package mymodel

import (
	"DEMOX_ADMINAUTH/internal/pkg/dbtype"
)

// my info
// @doc | mymodel.MyInfo
type MyInfo struct {
	dbtype.Base
	Account      string `json:"account" gorm:"type:varchar(50);unique;comment:账号" binding:"required" doc:"|d 账号"`
	Phone        string `json:"phone" gorm:"type:varchar(50);unique;comment:电话" binding:"max=11" doc:"|d 电话"`
	Name         string `json:"name" gorm:"size:100;not null;default:'';comment:名称" binding:"max=100" doc:"|d 姓名"`
	Memo         string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300" doc:"|d 备注"`
	Email        string `json:"email" gorm:"type:varchar(100);default:'';comment:邮件" binding:"omitempty,email" doc:"|d Email"`
	IsSuperAdmin string `json:"is_super_admin" gorm:"type:enum('1','0');default:'0';comment:是否超级管理员" binding:"oneof=0 1" doc:"|d 是否超级管理员 |c 0不是 1是"`
	Role         string `json:"role" gorm:"size:100;not null;index;comment:角色" doc:"|d 角色ID"` //角色代码
}

func (MyInfo) TableName() string {
	return "admin"
}

// my update
// @doc | mymodel.MyForm
type MyForm struct {
	//dbtype.Base
	ID    int64  `json:"id" gorm:"primaryKey"`
	Phone string `json:"phone" gorm:"type:varchar(50);unique;comment:电话" binding:"max=11" doc:"|d 电话 |e 12312312312"`
	Name  string `json:"name" gorm:"size:100;not null;default:'';comment:名称" binding:"max=100" doc:"|d 姓名 |e 张三"`
	Memo  string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300" doc:"|d 备注 |e 这是备注"`
	Email string `json:"email" gorm:"type:varchar(100);default:'';comment:邮件" binding:"omitempty,email" doc:"|d Email |e 2@qq.com"`
}

func (MyForm) TableName() string {
	return "admin"
}

// log form
// @doc | mymodel.LoginForm
type LoginForm struct {
	Account  string `json:"account" binging:"required" doc:"|d 账号 |e admin"`
	Password string `json:"password" binging:"required" doc:"|d 密码 |e 12345"`
}

// 登陆返回token
// @doc | mymodel.LogRep
type LogRep struct {
	AccessToken  string `json:"access_token" doc:"|d 鉴权token |c header头Authorization参数"`
	TokenExp     int64  `json:"token_exp" doc:"|d 刷新时间戳 |c token有效期"`
	RefreshToken string `json:"refresh_token" doc:"|d 刷新token |c 刷新token时所用参数"`
}

// 刷新用token
// @doc | mymodel.RefreshTokeForm
type RefreshTokeForm struct {
	RefreshToken string `json:"refresh_token" binding:"required" doc:"|d 刷新token"`
}

// 修改密码
// @doc | mymodel.PasswordForm
type PasswordForm struct {
	Password    string `json:"password" binding:"required" doc:"|d 原始密码"`
	NewPassword string `json:"new_password" binding:"required" doc:"|d 新密码"`
}
