package tracex

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	gtrace "go.opentelemetry.io/otel/trace"
	"sync"
	"sync/atomic"
	"time"
)

var traceIsRun atomic.Bool
var traceOnce sync.Once
var _trace *Trace

type Trace struct {
	lock        sync.Mutex
	config      Config
	Trace       gtrace.Tracer
	Provider    *trace.TracerProvider
	Propagation propagation.TextMapPropagator
}

func InitTrace(cfg Config) {
	if _trace == nil {
		traceOnce.Do(func() {
			_trace = &Trace{
				config: cfg,
			}
			traceIsRun.Store(false)
		})
	}
}

// 启动trace
func (this *Trace) Run() error {
	this.lock.Lock()
	defer this.lock.Unlock()

	if err := this.initTraceSdk(); err != nil {
		return err
	}
	traceIsRun.Store(true)

	return nil
}

func (this *Trace) initTraceSdk() error {

	//传播器
	this.Propagation = this.newPropagator()
	otel.SetTextMapPropagator(this.Propagation)
	log.Println("===========", this.config.EndpointUrl, this.config.Auth)
	export, err := otlptracehttp.New(context.Background(),
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithEndpointURL(this.config.EndpointUrl),
		otlptracehttp.WithHeaders(map[string]string{
			"Authorization": this.config.Auth,
			"stream-name":   this.config.StreamName,
		}))

	if err != nil {
		log.Printf("Failed to initialize trace exporter: %v", err)
		return err
	}

	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL, semconv.ServiceName(this.config.ServerName),
		),
	)
	if err != nil {
		log.Printf("Failed to merge resources: %v", err)
		return err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(export,
			// Default is 5s. Set to 1s for demonstrative purposes.
			trace.WithBatchTimeout(time.Second)),

		trace.WithResource(r),
	)
	this.Provider = traceProvider
	otel.SetTracerProvider(traceProvider)
	this.Trace = otel.Tracer(this.config.ServerName)
	//this.Trace = traceProvider.Tracer(this.config.ServerName)
	log.Println("Trace SDK initialized successfully")
	return nil
}

func (this *Trace) newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

// 停止trace
func (this *Trace) Stop() {
	this.lock.Lock()
	defer this.lock.Unlock()
	traceIsRun.Store(false)
	if this.Provider != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if err := this.Provider.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down trace provider: %v", err)
		} else {
			this.Provider = nil
			this.Trace = nil
			log.Printf("Trace provider shutdown successfully")
		}
	}
}

// // 封装 使用
func (this *Trace) Start(ctx context.Context, name string, opts ...gtrace.SpanStartOption) (context.Context, Spanx) {
	if this.Trace != nil && traceIsRun.Load() {
		ctx, s := this.Trace.Start(ctx, name, opts...)
		return ctx, Spanx{s}
	} else {
		return nil, Spanx{}
	}
}

// 启动
func Run() error {
	if _trace != nil {
		return _trace.Run()
	} else {
		errors.New("trace 未初始化")
	}
	return _trace.Run()
}

// 停止
func Stop() {
	if _trace != nil {
		_trace.Stop()
	}
}

func Start(ctx context.Context, name string, opts ...gtrace.SpanStartOption) (context.Context, Spanx) {
	return _trace.Start(ctx, name, opts...)
}

func GetTrace() gtrace.Tracer {
	return _trace.Trace
}

func SafeSpan(f func()) error {
	if traceIsRun.Load() {
		f()
		return nil
	} else {
		return errors.New("trace not run")
	}
}

func GetEable() bool {
	return traceIsRun.Load()
}
