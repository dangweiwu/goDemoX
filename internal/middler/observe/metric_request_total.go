package observe

import (
	"DEMOX_ADMINAUTH/internal/pkg/observe/metric"
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	gmetric "go.opentelemetry.io/otel/metric"
	"log"
)

// 指标采集 请求数 计算qps
func RequestTotal(metricname string) gin.HandlerFunc {
	mt := metric.GetMeter()
	if mt == nil {
		panic("metric not init")
	}
	//每次新建会有一个starttime，就算metricname相同，也不是同个对象。
	requestTotal, err := mt.Int64Counter(
		metricname,
		gmetric.WithUnit("By"),
		gmetric.WithDescription("api request total."),
	)
	if err != nil {
		log.Printf("MetricRequestTotal err %v", err)
	}

	return func(c *gin.Context) {
		if !metric.GetEnable() {
			c.Next()
			return
		}

		c.Next()

		requestTotal.Add(context.Background(), 1, gmetric.WithAttributes(
			attribute.String("path", c.Request.Method+":"+c.FullPath()),
		))
	}
}
