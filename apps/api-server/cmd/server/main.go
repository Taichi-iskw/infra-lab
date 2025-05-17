package main

import (
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

func main() {
	r := gin.Default()

	setupPingRouter(r)
	setupHealthzRouter(r)
	setupPrometheusRouter(r)

	r.Run(":8080")
}
