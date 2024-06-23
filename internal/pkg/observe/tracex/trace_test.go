package tracex

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"log"
	"testing"
	"time"
)

func PrintPtr(a interface{}) {
	fmt.Printf("ptr:%p\n", a)
}

type A struct {
	s string
}

func (this A) Getname() string {
	PrintPtr(&this)
	return this.s
}

func (this *A) SetName(a string) {
	PrintPtr(this)
	(*this).s = a
	//this.s = a
}

func TestT(t *testing.T) {
	return
	InitTrace(Config{
		ServerName:  "test",
		EndpointUrl: "http://192.168.3.30:5080/api/abc/traces",
		Auth:        "Basic cm9vdEBxcS5jb206VlpWMHc1akZDRDhSWm9rMA==",
		StreamName:  "abc",
	})
	log.Println("start trace")
	Run()
	Stop()
	Run()

	log.Println("run trace")

	ctx := context.Background()

	ctx, span := Start(ctx, "empty")

	span.SetAttributes(attribute.Bool("isTrue", true), attribute.String("stringAttr", "hi!"))
	time.Sleep(time.Second)
	span.AddEvent("Cancelled wait due to external signal", trace.WithAttributes(attribute.Int("pid", 4328), attribute.String("signal", "SIGHUP")))
	time.Sleep(time.Second)
	span.SetStatus(codes.Error, "operationThatCouldFail failed")
	span.RecordError(errors.New("over error"))
	log.Println("span2 start")
	ctx, span2 := Start(ctx, "trace name2")
	//
	span2.SetAttributes(attribute.Bool("2 isTrue", true), attribute.String("2 stringAttr", "hi!"))
	span2.AddEvent("2 Cancelled wait due to external signal", trace.WithAttributes(attribute.Int("2 pid", 4328), attribute.String("2 signal", "SIGHUP")))
	span2.SetStatus(codes.Error, "2 operationThatCouldFail failed")
	span2.RecordError(errors.New("2 over error"))
	time.Sleep(time.Second)
	span2.End()
	span.End()

	time.Sleep(time.Second * 5)
	log.Println("over")
}
