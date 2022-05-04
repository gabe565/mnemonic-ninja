package telemetry

import (
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

func init() {
	flag.String("telemetry-name", "mnemonic-ninja", "OpenTelemetry service name")
}

func NewTraceProvider(exp *otlptrace.Exporter) *sdktrace.TracerProvider {
	// The service.name attribute is required.
	//
	//
	// Your service name will be used as the Service Dataset in honeycomb, which is where data is stored.
	r :=
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(viper.GetString("telemetry.name")),
		)

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}
