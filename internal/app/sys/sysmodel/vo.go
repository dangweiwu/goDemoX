package sysmodel

// @doc | sysmodel.SysVo
type SysVo struct {
	RunTime    string `json:"run_time" doc:"|d 运行时间 |t string |c 日期格式化字符串"`
	StartTime  string `json:"start_time" doc:"|d 开始时间 |t string |c 日期格式化字符串"`
	OpenTrace  string `json:"open_trace" binding:"onof=0 1" doc:"|d 链路追踪开关"`
	OpenMetric string `json:"open_metric" binding:"oneof=0 1" doc:"|d 指标采集开关"`
}

func (SysVo) TableName() string {
	return "sys"
}
