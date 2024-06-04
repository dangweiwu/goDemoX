package ctx

import (
	"DEMOX_ADMINAUTH/internal/config"
	"DEMOX_ADMINAUTH/internal/pkg/log"
	"github.com/dangweiwu/ginpro/pkg/mysqlx"
	errs "github.com/pkg/errors"
)

func NewDbContext(c config.Config) (*AppContext, error) {
	//初始化日志
	appctx := &AppContext{}
	appctx.Config = c
	if lg, err := log.New(c.Log); err != nil {
		return nil, err
	} else {
		appctx.Log = lg
	}

	//初始化数据库
	db := mysqlx.NewDb(c.Mysql)
	if d, err := db.GetDb(); err != nil {
		return nil, errs.WithMessage(err, "err init db")
	} else {
		//d.Debug()
		appctx.Db = d
		log.Msg("数据库链接成功").Info(appctx.Log)
	}
	return appctx, nil
}
