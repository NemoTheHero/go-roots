package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexController(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello")

}
