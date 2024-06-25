package middler

import (
	"github.com/dangweiwu/ginpro/pkg/metric"
	"github.com/gin-gonic/gin"
	"goDemoX/internal/ctx"
	"strconv"
	"time"
)

//指标中间件

var skipMapPath = map[string]struct{}{}

func PromMiddler(appctx *ctx.AppContext) gin.HandlerFunc {
	var metricServerReqDur = metric.NewHistogramVec(&metric.HistogramVecOpts{
		Namespace: appctx.Config.App.Name,
		Subsystem: "requests",
		Name:      "duration_ms",
		Help:      "http server requests duration(ms).",
		Labels:    []string{"path"},
		Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	})
	//sum(rate(http_server_requests_duration_ms_count{app="http"}[3m])) by (path)
	//”http_server_requests_duration_ms“ bucket count sum
	var metricServerReqCodeTotal = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: appctx.Config.App.Name,
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "http server requests error count.",
		Labels:    []string{"path", "code"},
	})
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		if _, has := skipPath[path]; !has {
			startTime := time.Now()
			context.Next()
			metricServerReqDur.Observe(time.Now().UnixMilli()-startTime.UnixMilli(), path)
			metricServerReqCodeTotal.Inc(path, strconv.Itoa(context.Writer.Status()))
		}
	}
}
