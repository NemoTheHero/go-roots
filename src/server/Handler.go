package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"roots/src/controller"
	"roots/src/dao"
)

func Handler() {
	router := gin.Default()
	v1 := router.Group("/").Use(SecurityMiddleWare)
	{
		v1.GET("test", controller.TestController)
		v1.GET("", controller.IndexController)
		v1.GET("albums", controller.GetAlbums)
		v1.POST("albums", controller.PostAlbums)
		v1.GET("albums/:id", controller.GetAlbumByID)
		v1.POST("createTable", dao.AmazonDao)
		v1.POST("addItem", dao.AddItem)
		v1.GET("getMovies", dao.GetMovies)
		v1.GET("getAbove4", dao.GetAbove4)
		v1.PUT("updateMovie", dao.UpdateItem)
	}

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "port"), nil))

	}

}
