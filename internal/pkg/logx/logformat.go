package logx

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 日志格式化
const (
	DATA        = "data"
	DATAEX      = "dataex"
	JSON_FORMAT = "json"
	KIND        = "kind"
	TRACEID     = "traceid"
)

type Format struct {
	msg  string
	data []zapcore.Field
}

func Msg(msg ...string) *Format {
	m := ""
	if len(msg) > 0 {
		m = msg[0]
	}
	return &Format{
		data: []zapcore.Field{},
		msg:  m,
	}
}

func (this *Format) Msg(msg string) *Format {
	this.msg = msg
	return this
}

func (this *Format) FmtData(format string, args ...interface{}) *Format {
	this.data = append(this.data, zap.String(DATA, fmt.Sprintf(format, args...)))
	return this
}

func (this *Format) Data(data string) *Format {
	this.data = append(this.data, zap.String(DATA, data))
	return this
}

func (this *Format) DataEx(data string) *Format {
	this.data = append(this.data, zap.String(DATAEX, data))
	return this
}

func (this *Format) Kind(data string) *Format {
	this.data = append(this.data, zap.String(KIND, data))
	return this
}

func (this *Format) Trace(data string) *Format {
	this.data = append(this.data, zap.String(TRACEID, data))
	return this
}

func (this *Format) ErrData(err error) *Format {
	this.data = append(this.data, zap.Error(err))
	return this
}

func (this *Format) Json(data interface{}) {
	d, _ := json.Marshal(data)
	this.data = append(this.data, zap.String(JSON_FORMAT, string(d)))
}

func (this *Format) Info(lg *zap.Logger) {
	lg.Log(zapcore.InfoLevel, this.msg, this.data...)
}

func (this *Format) Err(lg *zap.Logger) {
	lg.Log(zapcore.ErrorLevel, this.msg, this.data...)
}

func (this *Format) Debug(lg *zap.Logger) {
	lg.Log(zapcore.DebugLevel, this.msg, this.data...)
}

func (this *Format) Panic(lg *zap.Logger) {
	lg.Log(zapcore.PanicLevel, this.msg, this.data...)
}
