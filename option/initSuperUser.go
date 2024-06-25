package option

import (
	"github.com/dangweiwu/ginpro/pkg/yamconfig"
	"goDemoX/internal/app/admin/adminmodel"
	"goDemoX/internal/config"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg"
	"gorm.io/gorm"
	"log"
)

type InitSuperUser struct {
	Password string `long:"password" description:"超级管理员设置密码"`
}

func (this *InitSuperUser) Usage() string {
	return `
设置超级管理员密码`
}

func (this *InitSuperUser) Execute(args []string) error {
	var c config.Config
	yamconfig.MustLoad(Opt.ConfigPath, &c)
	appctx, err := ctx.NewDbContext(c)
	if err != nil {
		panic(err)
	}
	po := &adminmodel.AdminPo{}
	pwd := this.Password
	if this.Password == "" {
		pwd = "a123456"
	}
	if r := appctx.Db.Where("account = 'admin'").Take(po); r.Error == gorm.ErrRecordNotFound {
		//创建
		po.Name = "超级管理员"
		po.Account = "admin"
		po.IsSuperAdmin = "1"
		po.Password = pkg.GetPassword(pwd)
		if r := appctx.Db.Create(po); r.Error != nil {
			return r.Error
		}
	} else if r.Error == nil {
		//更新
		po.Password = pkg.GetPassword(pwd)
		if r := appctx.Db.Select("password").Updates(po); r.Error != nil {
			return r.Error
		}
		log.Println("密码设置成功：", pwd)
	} else {
		return r.Error
	}
	return nil
}
