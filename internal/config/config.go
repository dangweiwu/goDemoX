package config

import (
	"goDemoX/internal/pkg/api/apiserver/apiconfig"
	"goDemoX/internal/pkg/db/mysqlx"
	"goDemoX/internal/pkg/db/redisx"
	"goDemoX/internal/pkg/jwtx/jwtconfig"
	"goDemoX/internal/pkg/logx"
)

// 全局配置文件
type Config struct {
	App    App
	Api    apiconfig.ApiConfig
	Log    logx.Config
	Mysql  mysqlx.Config
	Jwt    jwtconfig.JwtConfig
	Redis  redisx.Config
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
