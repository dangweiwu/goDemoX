package sysmodel

// @doc | sysmodel.SysActForm
type SysActForm struct {
	Name string `json:"name" binding:"oneof=trace metric" doc:"|d 名称"`
	Act  string `json:"act" binding:"oneof=0 1" doc:"|d 开关"` //0 启动 1 停止
}

func (SysActForm) TableName() string {
	return "sys"
}
