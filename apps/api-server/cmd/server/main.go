package main

import "github.com/gin-gonic/gin"

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

func main() {
	r := gin.Default()

	setupPingRouter(r)
	setupHealthzRouter(r)

	r.Run(":8080")
}
