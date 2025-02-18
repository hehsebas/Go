package main

import (
	"API-RESTful/config"
	"API-RESTful/controllers"
	"API-RESTful/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.InitMySQL(); err != nil {
		panic(err)
	}
	defer config.DB.Close()
	albumRepo := repository.NewAlbumRepository(config.DB)
	albumController := controllers.NewAlbumController(albumRepo)
	r := gin.Default()
	r.GET("/albums", albumController.ShowAlbums)
	r.POST("/albums", albumController.AddAlbum)
	r.DELETE("/albums/:id", albumController.DeleteAlbum)
	r.PUT("/albums/:id", albumController.ModifyAlbum)
	r.Run("localhost:8080")
}
