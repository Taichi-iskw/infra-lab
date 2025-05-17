package main

import (
	"math"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequests = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Total number of HTTP requests",
})

func init() {
	prometheus.MustRegister(httpRequests)
}

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
	r.GET("/metrics", func(c *gin.Context) {
		httpRequests.Inc()
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
		c.JSON(200, gin.H{"message": "ok"})
	})
}

func setupComputeRouter(r *gin.Engine) {
	r.GET("/compute", func(c *gin.Context) {
		// cpu intensive computation
		for i := range 1000000000 {
			math.Sqrt(float64(i))
		}

		// memory intensive computation
		memory := make([]byte, 1024*1024*1024)
		for i := range memory {
			memory[i] = byte(i)
		}

		c.JSON(200, gin.H{"message": "ok"})
	})
}

func main() {
	r := gin.Default()

	setupPingRouter(r)
	setupHealthzRouter(r)
	setupPrometheusRouter(r)
	setupComputeRouter(r)

	r.Run(":8080")
}
