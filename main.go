package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// now time
		t := time.Now()
		// handle the request
		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()
		// log latency and status  too
		log.Printf("latency: %s, status: %d", latency, status)
	}
}

func main() {
	r := gin.Default()
	r.Use(Logger())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
