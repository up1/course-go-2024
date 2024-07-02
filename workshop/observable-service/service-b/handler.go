package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
)

const name = "service-b"

var (
	tracer = otel.Tracer(name)
	meter  = otel.Meter(name)
	logger = otelslog.NewLogger(name)
)

func SetupHandler() *gin.Engine {
	router := gin.New()
	router.Use(otelgin.Middleware("my-server"))
	router.GET("/data", getData)
	return router
}

func getData(c *gin.Context) {
	// Trace
	_, span := tracer.Start(c.Request.Context(), "GetData")
	defer span.End()

	// Log
	logger.Info("Called /data", slog.Attr{Key: "traceid", Value: slog.StringValue(span.SpanContext().TraceID().String())})

	c.JSON(http.StatusOK, gin.H{
		"message": "hell from /data",
	})
}
