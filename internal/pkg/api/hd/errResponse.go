package hd

type ErrKind string

const (
	CODE ErrKind = "code" //需要映射 code映射成相关中文名 多语言需要
	MSG  ErrKind = "msg"  //无需映射 直接显示
)

// @doc | hd.ErrResponse
type ErrResponse struct {
	Kind ErrKind `json:"kind" doc:"|d 类型 |c code msg |t string"`
	Data string  `json:"data" doc:"|d 代码 |c kind=code时候 data指示异常代码 |t string"`
	Msg  string  `json:"msg" doc:"|d 消息 |c kind=msg时候 msg标识异常信息 |t string"`
}

func (this ErrResponse) Error() string {
	return this.Data
}

func ErrCode(data string, msg string) ErrResponse {
	return ErrResponse{CODE, data, msg}
}

func ErrMsg(data string, msg string) ErrResponse {
	return ErrResponse{MSG, data, msg}
}
