package metricx

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	gmetric "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"log"

	"sync"
	"sync/atomic"
	"time"
)

var metricIsRun atomic.Bool
var metricOnce sync.Once
var _metric *Metric

type Metric struct {
	lock     sync.Mutex
	config   Config
	Provider *metric.MeterProvider
	Meter    gmetric.Meter
}

func InitMetric(cfg Config) {
	if _metric == nil {
		metricOnce.Do(func() {
			_metric = &Metric{
				config: cfg,
				Meter:  otel.Meter(cfg.ServerName),
			}
			metricIsRun.Store(false)
			//_metric.Run()
		})
	}
}

// 初始化sdk
func (this *Metric) initSdk() error {

	metricExporter, err := otlpmetrichttp.New(context.Background(), otlpmetrichttp.WithInsecure(),
		otlpmetrichttp.WithEndpointURL(this.config.EndpointUrl),
		otlpmetrichttp.WithHeaders(map[string]string{
			"Authorization": this.config.Auth,
			"stream-name":   this.config.StreamName,
		}))
	if err != nil {
		log.Printf("Failed to initialize trace exporter: %v", err)
		return err
	}
	//resource.Default()
	//r, err := resource.Merge(
	//	resource.Empty(),
	//	//resource.Default(),
	//	resource.NewWithAttributes(
	//		semconv.SchemaURL, semconv.ServiceName(this.config.ServerName),
	//	),
	//)
	//if err != nil {
	//	log.Println("resource.Merge err", err)
	//	return err
	//}
	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter,
			// Default is 1m. Set to 3s for demonstrative purposes.
			metric.WithInterval(time.Duration(this.config.Interval)*time.Second)),
		),
		metric.WithResource(resource.Empty()),
	)
	this.Provider = meterProvider

	otel.SetMeterProvider(this.Provider)
	//this.Meter = meterProvider.Meter(this.config.ServerName)
	log.Println("Metric SDK initialized successfully")
	return nil
}

func (this *Metric) Run() {
	this.lock.Lock()
	defer this.lock.Unlock()
	if !metricIsRun.Load() {
		err := this.initSdk()
		if err != nil {
			log.Println("initSdk err", err)
			return
		}
		metricIsRun.Store(true)
	}
}

func (this *Metric) Stop() {
	this.lock.Lock()
	defer this.lock.Unlock()
	if metricIsRun.Load() {
		this.Provider.Shutdown(context.Background())
		metricIsRun.Store(false)
	}
}

func Run() {
	if _metric != nil {
		_metric.Run()
	}
}

func Stop() {
	if _metric != nil {
		_metric.Stop()
	}
}

func GetEnable() bool {
	return metricIsRun.Load()
}

func GetMeter() gmetric.Meter {
	return _metric.Meter
}

func GetMertric() *Metric {
	return _metric
}
