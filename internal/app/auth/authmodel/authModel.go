package authmodel

import "goDemoX/internal/pkg/dbtype"

type AuthPo struct {
	dbtype.Base
	Name     string   `json:"name" gorm:"size:100;comment:名称" binding:"max=100" extensions:"x-name=名称,x-value="`
	Code     string   `json:"code" gorm:"size:100;not null;unique;comment:权限ID" binding:"required,max=100" extensions:"x-name=账号,x-value="`
	OrderNum int      `json:"order_num" gorm:"default:0;comment:排序" binding:"" extensions:"x-name=排序,x-value="`
	Api      string   `json:"api" gorm:"size:200;default:'';comment:接口" binding:"max=200" extensions:"x-name=API,x-value="`
	Method   string   `json:"method" gorm:"size:10;default:'';comment:请求方式" binding:"max=50" extensions:"x-name=方法,x-value="`
	Kind     string   `json:"kind" gorm:"type:enum('1','0');default:'0';comment:0 api 1 菜单" binding:"oneof=0 1" extensions:"x-name=类型,x-value="` //0 api 1 菜单
	ParentId int      `json:"parent_id" gorm:"default:0;not null;comment:上级菜单ID" extensions:"x-name=父ID,x-value="`
	Children []AuthPo `json:"children" gorm:"foreignkey:ParentId"`
}

func (AuthPo) TableName() string {
	return "auth"
}

// @doc | authmodel.AuthForm
type AuthForm struct {
	dbtype.BaseForm
	Name     string `json:"name" gorm:"size:100;comment:名称" binding:"max=100,required" doc:"|d 权限名称"`
	Code     string `json:"code" gorm:"size:100;not null;unique;comment:权限ID"  binding:"required,max=100" doc:"|d 权限ID"`
	OrderNum int    `json:"order_num" gorm:"default:0;comment:排序" binding:""  doc:"|d 排序"`
	Api      string `json:"api" gorm:"size:200;default:'';comment:接口" binding:"max=200" doc:"|d API"`
	Method   string `json:"method" gorm:"size:10;default:'';comment:请求方式" binding:"max=50"  doc:"|d 方法 |e GET"`                           // GET POST PUT DELETE
	Kind     string `json:"kind" gorm:"type:enum('1','0');default:'0';comment:0 api 1 菜单" binding:"oneof=0 1"  doc:"|d 类型 |c 0:api 1:菜单"` //0 api 1 菜单
	ParentId int    `json:"parent_id" gorm:"default:0;not null;comment:上级菜单ID"  doc:"|d 上级ID"`                                            //父类ID
}

func (AuthForm) TableName() string {
	return "auth"
}

// @doc| authmodel.AuthUpdateForm
type AuthUpdateForm struct {
	dbtype.BaseForm
	Name     string `json:"name" gorm:"size:100;comment:名称" binding:"max=100,required" doc:"|d 权限名称"`
	OrderNum int    `json:"order_num" gorm:"default:0;comment:排序" binding:""  doc:"|d 排序"`
	Api      string `json:"api" gorm:"size:200;default:'';comment:接口" binding:"max=200" doc:"|d API"`
	Method   string `json:"method" gorm:"size:10;default:'';comment:请求方式" binding:"max=50"  doc:"|d 方法 |e GET"`                           // GET POST PUT DELETE
	Kind     string `json:"kind" gorm:"type:enum('1','0');default:'0';comment:0 api 1 菜单" binding:"oneof=0 1"  doc:"|d 类型 |c 0:api 1:菜单"` //0 api 1 菜单
	ParentId int    `json:"parent_id" gorm:"default:0;not null;comment:上级菜单ID"  doc:"|d 上级ID"`                                            //父类ID
}

func (AuthUpdateForm) TableName() string {
	return "auth"
}

// @doc| authmodel.AuthVo
type AuthVo struct {
	dbtype.Base
	Name     string   `json:"name" gorm:"size:100;comment:名称" binding:"max=100,required" doc:"|d 权限名称"`
	Code     string   `json:"code" gorm:"size:100;not null;unique;comment:权限ID"  binding:"required,max=100" doc:"|d 权限ID"`
	OrderNum int      `json:"order_num" gorm:"default:0;comment:排序" binding:""  doc:"|d 排序"`
	Api      string   `json:"api" gorm:"size:200;default:'';comment:接口" binding:"max=200" doc:"|d API"`
	Method   string   `json:"method" gorm:"size:10;default:'';comment:请求方式" binding:"max=50"  doc:"|d 方法 |e GET"`                           // GET POST PUT DELETE
	Kind     string   `json:"kind" gorm:"type:enum('1','0');default:'0';comment:0 api 1 菜单" binding:"oneof=0 1"  doc:"|d 类型 |c 0:api 1:菜单"` //0 api 1 菜单
	ParentId int      `json:"parent_id" gorm:"default:0;not null;comment:上级菜单ID"  doc:"|d 上级ID"`
	Children []AuthVo `json:"children" gorm:"foreignkey:ParentId" doc:"|d 子集 |t []self "`
}

func (AuthVo) TableName() string {
	return "auth"
}
