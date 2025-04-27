package build

import (
	"context"

	"github.com/cockroachdb/errors"
	sentryotel "github.com/getsentry/sentry-go/otel"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func (b *Builder) tracer(ctx context.Context) error {
	if !b.config.UseTrace() {
		return nil
	}

	options := []otlptracegrpc.Option{otlptracegrpc.WithEndpoint(b.config.OTLPTrace.Endpoint)}
	if !b.config.OTLPTrace.TLS {
		options = append(options, otlptracegrpc.WithInsecure())
	}

	exporter, err := otlptracegrpc.New(ctx, options...)
	if err != nil {
		return errors.Wrap(err, "build trace exporter")
	}

	tp := trace.NewTracerProvider(
		trace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
		trace.WithSampler(trace.TraceIDRatioBased(b.config.App.TraceSampleRatio)),
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String(b.config.App.Name))),
	)

	b.shutdown.add(tp.Shutdown)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return nil
}
