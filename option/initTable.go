package option

import (
	"DEMOX_ADMINAUTH/internal/app"
	"DEMOX_ADMINAUTH/internal/config"
	"DEMOX_ADMINAUTH/internal/ctx"
	"fmt"
	"github.com/dangweiwu/ginpro/pkg/yamconfig"
	"log"
)

// 初始化数据库
type InitTable struct {
}

func (*InitTable) Usage() string {
	return `
  迁移数据库结构，但不会删除未使用的列。`
}
func (this *InitTable) Execute(args []string) error {

	var c config.Config
	yamconfig.MustLoad(Opt.ConfigPath, &c)
	fmt.Println(c.Mysql.Host)
	appctx, err := ctx.NewDbContext(c)
	if err != nil {
		panic(err)
	}

	app.Regdb(appctx)

	log.Println("初始化完毕")

	return nil
}
