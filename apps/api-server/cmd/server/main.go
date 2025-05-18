package main

import (
	"io"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// --- Prometheus Metrics ---

var httpRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"path", "status"},
)

var httpDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"path"},
)

func init() {
	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(httpDuration)
}

// --- Prometheus Middleware ---

func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()

		path := c.FullPath()
		if path == "" {
			path = "unknown"
		}
		status := strconv.Itoa(c.Writer.Status())

		httpRequests.WithLabelValues(path, status).Inc()
		httpDuration.WithLabelValues(path).Observe(duration)
	}
}

// --- Routers ---

func setupPingRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}

func setupHealthzRouter(r *gin.Engine) {
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})
}

func setupPrometheusRouter(r *gin.Engine) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

func setupComputeRouter(r *gin.Engine) {
	r.GET("/compute", func(c *gin.Context) {
		for i := range 1000000000 {
			math.Sqrt(float64(i))
		}
		memory := make([]byte, 1024*1024*1024)
		for i := range memory {
			memory[i] = byte(i)
		}
		c.JSON(200, gin.H{"message": "ok"})
	})
}

func setupLoadRouter(r *gin.Engine) {
	isTest := os.Getenv("IS_TEST")

	r.GET("/load", func(c *gin.Context) {
		for range 10 {
			if isTest == "true" {
				time.Sleep(1 * time.Millisecond)
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
		c.JSON(200, gin.H{"message": "ok"})
	})
}

func setupErrorRouter(r *gin.Engine) {
	r.GET("/error", func(c *gin.Context) {
		c.JSON(500, gin.H{"error": "simulated internal server error"})
	})
}

func setupEchoRouter(r *gin.Engine) {
	r.POST("/echo", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": "failed to read request body"})
			return
		}
		c.JSON(200, gin.H{"message": "ok", "body": string(body)})
	})
}

// --- Main ---

func main() {
	r := gin.Default()
	r.Use(prometheusMiddleware())

	setupPingRouter(r)
	setupHealthzRouter(r)
	setupPrometheusRouter(r)
	setupComputeRouter(r)
	setupLoadRouter(r)
	setupErrorRouter(r)
	setupEchoRouter(r)

	r.Run(":8080")
}
