package testtool

import (
	"DEMOX_ADMINAUTH/internal/config"
	"DEMOX_ADMINAUTH/internal/pkg/jwtx/jwtconfig"
	"DEMOX_ADMINAUTH/internal/pkg/log"
	"github.com/dangweiwu/ginpro/pkg/mysqlx/mysqlxconfig"
	"github.com/dangweiwu/ginpro/pkg/redisx/redisconfig"
)

func NewTestConfig() config.Config {
	a := config.Config{}
	a.App = config.App{Name: "test"}
	a.Log = log.Config{Level: "error", OutType: "console", Formatter: "json"}
	a.Redis = redisconfig.RedisConfig{}
	a.Mysql = mysqlxconfig.Mysql{Host: "localhost:4417", DbName: "test"}
	a.Jwt = jwtconfig.JwtConfig{"123", int64(5)}
	return a
}
