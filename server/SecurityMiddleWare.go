package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func SecurityMiddleWare(c *gin.Context) {
	t := time.Now()
	log.Println("before request")
	// Set example variable
	c.Set("example", "12345")

	// before request

	c.Next()

	// after request
	latency := time.Since(t)
	log.Print(latency)

	// access the status we are sending
	status := c.Writer.Status()
	log.Println(status)
	log.Println("security middleware")
}
