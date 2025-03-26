package controllers

import (
	"marketplace/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeHandler maneja la ruta raíz de la aplicación
func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Bienvenido a la API de Marketplace"})
}

// PingHandler maneja la ruta de prueba "/ping"
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Pong"})
}

// CrearUsuario maneja la creación de un nuevo usuario
func CrearUsuario(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := models.CrearUsuario(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado exitosamente", "user": user})
}

// ObtenerUsuarios maneja la obtención de todos los usuarios
func ObtenerUsuarios(c *gin.Context) {
	users, err := models.ObtenerUsuarios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// ObtenerUsuario maneja la obtención de un solo usuario por ID
func ObtenerUsuario(c *gin.Context) {
	id := c.Param("id")
	user, err := models.ObtenerUsuarioPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// ActualizarUsuario maneja la actualización de un usuario por ID
func ActualizarUsuario(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := models.ActualizarUsuario(id, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente", "user": user})
}

// EliminarUsuario maneja la eliminación de un usuario por ID
func EliminarUsuario(c *gin.Context) {
	id := c.Param("id")
	if err := models.EliminarUsuario(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente"})
}
