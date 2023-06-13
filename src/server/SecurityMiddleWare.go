package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"roots/src/services"
	"strings"
	"time"
)

func SecurityMiddleWare(c *gin.Context) {
	t := time.Now()
	log.Println("before request")
	// Set example variable
	c.Set("example", "12345")

	// before request
	//if not valid
	authHeader := c.GetHeader("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	jwtToken := splitToken[1]
	if services.VerifyJwt(jwtToken) {
		c.IndentedJSON(http.StatusUnauthorized, authHeader)
		c.Abort()
	}
	c.Next()

	// after request
	latency := time.Since(t)
	log.Print(latency)

	// access the status we are sending
	status := c.Writer.Status()
	log.Println(status)
	log.Println("security middleware")
}
