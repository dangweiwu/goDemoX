package metric

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	gmetric "go.opentelemetry.io/otel/metric"
	"math/rand"
	"testing"
	"time"
)

func TestMetric(t *testing.T) {

	//gm := otel.Meter("api")

	//requestTotal, err := gm.Int64Counter(
	//	"count_one",
	//	gmetric.WithUnit("By"),
	//	gmetric.WithDescription("api request total."),
	//)

	//requestTotal2, err := gm.Int64Counter(
	//	"count_one",
	//	gmetric.WithUnit("By"),
	//	gmetric.WithDescription("api request total."),
	//)
	//fmt.Println(err)

	bm := otel.Meter("api")
	//requestTotal2, _ := gm.Int64Counter(
	//	"how_two",
	//	gmetric.WithUnit("By"),
	//	gmetric.WithDescription("api request total."),
	//)

	requestDuration, err := bm.Int64Histogram("http_duration",
		gmetric.WithUnit("ms"),
		gmetric.WithDescription("api duration"),
		gmetric.WithExplicitBucketBoundaries(10, 50, 100, 200, 500, 800, 1000),
	)
	fmt.Println(err)
	InitMetric(Config{
		EndpointUrl: "http://127.0.0.1:5080/api/default/v1/metrics",
		Auth:        "Basic cm9vdEBxcS5jb206VlpWMHc1akZDRDhSWm9rMA==",
		ServerName:  "linglan",
		StreamName:  "default",
		Interval:    5, //导出时间间隔 单位秒
	})

	mt := GetMertric()
	mt.Run()

	fmt.Println(err)
	//mt.Stop()
	//mt.Run()
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子

	for {

		//requestTotal.Add(context.Background(), 1)
		fmt.Println("add 0")
		//time.Sleep(time.Second)
		//requestTotal2.Add(context.Background(), 1)
		//
		//fmt.Println("add 1")

		randomMilliseconds := rand.Int63n(1000)
		requestDuration.Record(context.Background(), randomMilliseconds)

		time.Sleep(time.Millisecond * 200)
	}

}
