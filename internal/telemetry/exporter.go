package telemetry

import (
	"context"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"google.golang.org/grpc/credentials"
)

func init() {
	flag.String("telemetry-endpoint", "", "OpenTelemetry endpoint")
	flag.StringToString("telemetry-headers", map[string]string{}, "OpenTelemetry headers")
}

func NewExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	// Configuration to export data to Honeycomb:
	//
	// 1. The Honeycomb endpoint
	// 2. Your API key, set as the x-honeycomb-team header
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(viper.GetString("telemetry.endpoint")),
		otlptracegrpc.WithHeaders(viper.GetStringMapString("telemetry.headers")),
		otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, "")),
	}

	client := otlptracegrpc.NewClient(opts...)
	return otlptrace.New(ctx, client)
}
