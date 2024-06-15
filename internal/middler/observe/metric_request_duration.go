package observe

import (
	"DEMOX_ADMINAUTH/internal/pkg/observe/metricx"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	gmetric "go.opentelemetry.io/otel/metric"
	"log"
	"time"
)

/*
*
采集请求时长
*/
func RequestDuration(metricname string) gin.HandlerFunc {
	mt := metricx.GetMeter()
	//每次新建会有一个starttime，就算metricname相同，也不是同个对象。
	requestDuration, err := mt.Float64Histogram(
		metricname,
		gmetric.WithUnit("ms"),
		gmetric.WithDescription("api request duration."),
	)
	if err != nil {
		log.Println(fmt.Sprintf("MetricRequestDuration err %v", err))
	}

	return func(c *gin.Context) {
		if !metricx.GetEnable() {
			c.Next()
			return
		}
		requestStartTime := time.Now()
		c.Next()
		elapsedTime := float64(time.Since(requestStartTime)) / float64(time.Millisecond)
		requestDuration.Record(context.Background(), elapsedTime, gmetric.WithAttributes(
			attribute.String("path", c.Request.Method+":"+c.FullPath()),
		))
	}
}
