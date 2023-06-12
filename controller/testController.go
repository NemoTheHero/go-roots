package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestController(c *gin.Context) {
	println("testcontroller")
	c.IndentedJSON(http.StatusOK, "Test")

}
