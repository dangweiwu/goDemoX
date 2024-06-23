package ctx

import (
	"DEMOX_ADMINAUTH/internal/config"
	"DEMOX_ADMINAUTH/internal/pkg/db/mysqlx"
	"DEMOX_ADMINAUTH/internal/pkg/logx"
	errs "github.com/pkg/errors"
)

func NewDbContext(c config.Config) (*AppContext, error) {
	//初始化日志
	appctx := &AppContext{}
	appctx.Config = c
	if lg, err := logx.New(c.Log); err != nil {
		return nil, err
	} else {
		appctx.Log = lg
	}

	//初始化数据库
	if db, err := mysqlx.NewDb(c.Mysql); err != nil {
		return nil, errs.WithMessage(err, "err init db")
	} else {
		//d.Debug()
		appctx.Db = db.GetDb()
		logx.Msg("数据库链接成功").Info(appctx.Log)
	}
	return appctx, nil
}
