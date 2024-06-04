package app

import (
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/app/auth/authmodel"
	"DEMOX_ADMINAUTH/internal/app/role/rolemodel"
	"DEMOX_ADMINAUTH/internal/ctx"
)

var Tables = []interface{}{
	&adminmodel.AdminPo{},
	&authmodel.AuthPo{},
	&rolemodel.RolePo{},
}

func Regdb(appctx *ctx.AppContext) error {
	return appctx.Db.Set("gorm:ble_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(Tables...)
}
