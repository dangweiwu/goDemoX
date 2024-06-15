package config

import (
	"DEMOX_ADMINAUTH/internal/pkg/api/apiserver/apiconfig"
	"DEMOX_ADMINAUTH/internal/pkg/jwtx/jwtconfig"
	"DEMOX_ADMINAUTH/internal/pkg/logx"
	"github.com/dangweiwu/ginpro/pkg/mysqlx/mysqlxconfig"
	"github.com/dangweiwu/ginpro/pkg/redisx/redisconfig"
)

// 全局配置文件
type Config struct {
	App    App
	Api    apiconfig.ApiConfig
	Log    logx.Config
	Mysql  mysqlxconfig.Mysql
	Jwt    jwtconfig.JwtConfig
	Redis  redisconfig.RedisConfig
	Trace  TraceCfg
	Metric MetricCfg
}

// app
type App struct {
	Name     string
	Password string
}

// trace
type TraceCfg struct {
	Enable      bool
	EndpointUrl string // 链路追踪地址
	Auth        string // 链路追踪认证
	ServerName  string // 服务名称
	StreamName  string // 流名称
}

type MetricCfg struct {
	Enable      bool
	EndpointUrl string // 链路追踪地址
	Auth        string // 链路追踪认证
	ServerName  string // 服务名称
	StreamName  string
	Interval    int //导出时间间隔 单位秒
}
