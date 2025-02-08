package controllers

import (
	"API-RESTful/models"
	"API-RESTful/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	AlbumRepo *repository.AlbumRepository
}

// NewAlbumController crea una nueva instancia de AlbumController
func NewAlbumController(albumRepo *repository.AlbumRepository) *AlbumController {
	return &AlbumController{AlbumRepo: albumRepo}
}

// ShowAlbums obtiene y muestra todos los álbumes
func (c *AlbumController) ShowAlbums(ctx *gin.Context) {
	albums, err := c.AlbumRepo.GetAlbums()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, albums)
}

// AddAlbum agrega un nuevo álbum
func (c *AlbumController) AddAlbum(ctx *gin.Context) {
	var album models.Album
	if err := ctx.ShouldBindJSON(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.AlbumRepo.AddAlbum(album); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, album)
}
