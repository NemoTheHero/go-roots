package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"roots/controller"
)

func Handler() {
	router := gin.Default()
	v1 := router.Group("/v1").Use(SecurityMiddleWare)
	{
		v1.GET("/test", controller.TestController)
		v1.GET("/", controller.IndexController)
		v1.GET("/albums", controller.GetAlbums)
		v1.POST("/albums", controller.PostAlbums)
		v1.GET("/albums/:id", controller.GetAlbumByID)
	}

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "port"), nil))

	}

}
