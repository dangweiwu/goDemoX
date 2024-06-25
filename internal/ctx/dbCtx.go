package ctx

import (
	errs "github.com/pkg/errors"
	"goDemoX/internal/config"
	"goDemoX/internal/pkg/db/mysqlx"
	"goDemoX/internal/pkg/logx"
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
