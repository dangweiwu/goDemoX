package rolemodel

import (
	"goDemoX/internal/pkg/dbtype"
)

// @doc|rolemodel.RolePo
type RolePo struct {
	dbtype.Base
	Code     string              `json:"code" gorm:"size:100;not null;unique;comment:角色ID" binding:"required,max=100" doc:"|d 编码"`
	Name     string              `json:"name" gorm:"size:100;comment:角色名称" binding:"max=100" doc:"|d 名称"`
	OrderNum int                 `json:"order_num" gorm:"default:0;comment:排序" doc:"|d 排序"`
	Status   string              `json:"status" gorm:"type:enum('0','1');default:'1';comment:状态" doc:"|d 状态 |c 0:禁用 1:启用"` //0 禁用 1启用
	Memo     string              `json:"memo" gorm:"type:text;comment:备注" binding:"max=300" doc:"|d 备注"`
	Auth     dbtype.List[string] `json:"auth" gorm:"type:text;comment:角色code列表" doc:"|d 权限ID |t []string"` //权限编码列表 eg [auth1,...]
}

func (RolePo) TableName() string {
	return "role"
}

// @doc | rolemodel.RoleForm
type RoleForm struct {
	dbtype.BaseForm
	Code     string `json:"code" gorm:"size:100;not null;unique;comment:角色ID" binding:"required,max=100" doc:"|d 编码"` //全服唯一，禁止重复
	Name     string `json:"name" gorm:"size:100;comment:角色名称" binding:"max=100" doc:"|d 名称"`
	OrderNum int    `json:"order_num" gorm:"default:0;comment:排序" doc:"|d 权限代码 |c 6位编码12顶级菜单34当前菜单56接口编码"` //建议6位编码12顶级菜单34当前菜单56接口编码
	Status   string `json:"status" gorm:"type:enum('0','1');default:'1';comment:状态" doc:"|d 状态"`           //1 启动 0禁用
	Memo     string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300" doc:"|d 备注"`
}

func (RoleForm) TableName() string {
	return "role"
}

// @doc | rolemodel.RoleUpdate
type RoleUpdate struct {
	dbtype.BaseForm
	Name     string `json:"name" gorm:"size:100;comment:角色名称" binding:"max=100" doc:"|d 名称"`
	OrderNum int    `json:"order_num" gorm:"default:0;comment:排序" doc:"|d 权限代码 |c 6位编码12顶级菜单34当前菜单56接口编码"`   //建议6位编码12顶级菜单34当前菜单56接口编码
	Status   string `json:"status" gorm:"type:enum('0','1');default:'1';comment:状态" doc:"|d 状态 |c 0:禁用1:启用"` //1 启动 0禁用
	Memo     string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300" doc:"|d 备注"`
}

func (RoleUpdate) TableName() string {
	return "role"
}

// @doc | rolemodel.RoleAuthForm
type RoleAuthForm struct {
	ID   int64               `json:"id" swaggerignore:"true" gorm:"primaryKey"`
	Auth dbtype.List[string] `json:"auth" gorm:"type:text;comment:角色code列表" doc:"|d 权限列表 |t []string |e [auth1,auth2...]"` //权限编码列表 eg [auth1,auth2...]
}

func (RoleAuthForm) TableName() string {
	return "role"
}
