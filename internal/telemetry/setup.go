package telemetry

import (
	"context"
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"
)

func init() {
	flag.Bool("telemetry-enabled", false, "Enables OpenTelemetry")
}

func Setup() (*trace.TracerProvider, error) {
	if !viper.GetBool("telemetry.enabled") {
		return nil, nil
	}

	log.Println("Initialize OpenTelemetry")

	// Configure a new exporter using environment variables for sending data to Honeycomb over gRPC.
	exp, err := NewExporter(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize exporter: %v", err)
	}

	// Create a new tracer provider with a batch span processor and the otlp exporter.
	tp := NewTraceProvider(exp)

	// Set the Tracer Provider and the W3C Trace Context propagator as globals
	otel.SetTracerProvider(tp)

	// Register the trace context and baggage propagators so data is propagated across services/processes.
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return tp, nil
}
