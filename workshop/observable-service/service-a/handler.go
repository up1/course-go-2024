package api

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

func SetupHandler() *gin.Engine {
	router := gin.New()
	router.Use(otelgin.Middleware("my-server"))
	router.GET("/health", checkHealth)
	return router
}

const name = "service-a"

var (
	tracer      = otel.Tracer(name)
	meter       = otel.Meter(name)
	logger      = otelslog.NewLogger(name)
	calledCount metric.Int64Counter
)

func checkHealth(c *gin.Context) {
	// Metric
	calledCount, _ = meter.Int64Counter("health.counter",
		metric.WithDescription("The number of called health"),
		metric.WithUnit("{counts}"))
	calledCount.Add(c, 1)

	// Trace
	_, span := tracer.Start(c.Request.Context(), "CheckHealth")
	defer span.End()

	// Log
	logger.Info("Called /health", slog.Attr{Key: "traceid", Value: slog.StringValue(span.SpanContext().TraceID().String())})

	// Call service B
	ctx := trace.ContextWithSpan(c.Request.Context(), span)
	callServiceB(span, ctx)

	c.JSON(http.StatusOK, gin.H{
		"message": "healthy",
	})
}

func callServiceB(currentSpan trace.Span, ctx context.Context) {
	url := "http://service-b:8080/data"
	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	// Create new span from the parent span
	tr := currentSpan.TracerProvider().Tracer("service-a")
	ctx, span := tr.Start(ctx, "callServiceB", trace.WithAttributes(semconv.HTTPURLKey.String(url)))
	defer span.End()
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

	fmt.Printf("Sending request...\n")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	_, err = io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
}
