package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/contrib/propagators/jaeger"
	"go.opentelemetry.io/contrib/propagators/ot"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	argp := os.Args[1]
	argu := os.Args[2]
	start := time.Now()

	tracer := sdktrace.NewTracerProvider().Tracer("xk6/http")
	p := pickPropagator(argp)

	// Create the root span (aka. kickstart a new trace)
	ctx, span := tracer.Start(context.Background(), "")
	span.End()

	h := http.Header{}
	p.Inject(ctx, propagation.HeaderCarrier(h))

	elapsed := time.Since(start)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", argu, nil)
	req.Header = h
	resp, _ := client.Do(req)

	fmt.Println(resp.Status, span.SpanContext().TraceID().String(), elapsed, time.Since(start))
	fmt.Println(h)
}

func pickPropagator(propagator string) propagation.TextMapPropagator {
	switch propagator {
	case "w3c":
		return propagation.TraceContext{}
	case "b3":
		return b3.New()
	case "jaeger":
		return jaeger.Jaeger{}
	case "ot":
		return ot.OT{}
	case "aws":
		return xray.Propagator{}
	default:
		return propagation.TraceContext{}
	}
}
