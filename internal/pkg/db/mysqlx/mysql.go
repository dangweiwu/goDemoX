package mysqlx

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

//
//var _db *gorm.DB
//var once sync.Once

type Mysqlx struct {
	cfg Config
	_db *gorm.DB
}

func NewDb(cfg Config) (*Mysqlx, error) {

	dsn := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var out io.Writer
	if len(cfg.LogFile) != 0 {
		file, err := os.OpenFile(cfg.LogFile,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if err != nil {
			panic(err)
		}
		out = io.MultiWriter(file, os.Stdout)
	} else {
		out = os.Stdout
	}

	dbcfg := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.New(log.New(out, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
			logger.Config{
				SlowThreshold:             time.Second,                   // 慢 SQL 阈值
				LogLevel:                  logger.LogLevel(cfg.LogLevel), // 日志级别
				IgnoreRecordNotFoundError: false,                         // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,                         // 禁用彩色打印
			}),
	}
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(dsn, cfg.User, cfg.Password, cfg.Host, cfg.DbName)), dbcfg)
	return &Mysqlx{cfg: cfg, _db: db}, err
}

func (this *Mysqlx) GetDb() (db *gorm.DB) {
	return this._db
}
