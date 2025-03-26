package main

import (
	"log"
	"movilGO/config"
	"movilGO/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la conexión a MySQL
	config.InitMySQL()

	// Crear el router de Gin
	r := gin.Default()

	// Definir las rutas
	r.GET("/hello", controllers.GetHello)
	r.POST("/login", controllers.Login)

	// Iniciar el servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
