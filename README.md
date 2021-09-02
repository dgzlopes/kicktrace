# kicktrace

## HTTP



### w3c
Docs: https://www.w3.org/TR/baggage/#baggage-http-header-format
```
> go run main.go w3c https://test-api.k6.io
200 OK d439ca925f7d893bdda5db9ad41bfb11 73.151µs 443.248791ms
map[Traceparent:[00-d439ca925f7d893bdda5db9ad41bfb11-f591e7677390360d-01]]
```

### b3
Docs: https://github.com/openzipkin/b3-propagation
```
> go run main.go b3 https://test-api.k6.io
200 OK e5b44267e75bee7ba6c06a8a3c579980 70.961µs 473.51734ms
map[B3:[e5b44267e75bee7ba6c06a8a3c579980-bf809131175202b9-1]]
```

### jaeger
Docs: https://www.jaegertracing.io/docs/1.25/client-libraries/#propagation-format
```
> go run main.go jaeger https://test-api.k6.io
200 OK 78b51b4da5d18829b9bf576eb0f84ef0 44.451µs 505.419301ms
map[Uber-Trace-Id:[78b51b4da5d18829b9bf576eb0f84ef0:77198b3f71edfd17:0:1]]
```

### ot
Docs: https://opentracing.io/docs/overview/inject-extract/#required-carriers ??
```
> go run main.go ot https://test-api.k6.io
200 OK efc79732dac3928e28fc3191214f1696 43.931µs 442.294289ms
map[Ot-Tracer-Sampled:[1] Ot-Tracer-Spanid:[136b6ec28f79c086] Ot-Tracer-Traceid:[28fc3191214f1696]]
```

### aws/xray
Docs: ??
```
> go run main.go aws https://test-api.k6.io
200 OK c16a81d38aed9f4954462694fb2d78ea 62.271µs 474.730791ms
map[X-Amzn-Trace-Id:[Root=1-c16a81d3-8aed9f4954462694fb2d78ea;Parent=b60eb5867ab03835;Sampled=1]]
```

## gRPC

needs reseach -> https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/google.golang.org/grpc/otelgrpc