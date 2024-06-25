package app

import (
	"goDemoX/internal/app/admin/adminmodel"
	"goDemoX/internal/app/auth/authmodel"
	"goDemoX/internal/app/role/rolemodel"
	"goDemoX/internal/ctx"
)

var Tables = []interface{}{
	&adminmodel.AdminPo{},
	&authmodel.AuthPo{},
	&rolemodel.RolePo{},
}

func Regdb(appctx *ctx.AppContext) error {
	return appctx.Db.Set("gorm:ble_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(Tables...)
}
