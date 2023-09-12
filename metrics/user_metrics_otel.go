package metrics

import (
	"go.opentelemetry.io/otel"
	"log"

	"go.opentelemetry.io/otel/metric"
)

var (
	CreateUserCounter  metric.Int64Counter
	CreateUserDuration metric.Float64Histogram
)

func initMetrics(meter metric.Meter) {
	var err error
	// 注册测量工具
	CreateUserCounter, err = meter.Int64Counter("request_counter", metric.WithDescription("Request count"))
	if err != nil {
		log.Fatal("Error creating request counter:", err)
	}
	CreateUserDuration, err = meter.Float64Histogram("request_duration", metric.WithDescription("Request duration"))
	if err != nil {
		log.Fatal("Error creating request duration metric:", err)
	}
}

func init() {
	meter := otel.GetMeterProvider().Meter("user_data")
	initMetrics(meter)
}
