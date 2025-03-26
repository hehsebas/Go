package routes

import (
	"marketplace/controllers" // Importamos los controladores de tu proyecto

	"github.com/gin-gonic/gin" // Framework Gin
)

// RegisterRoutes configura todas las rutas de la aplicación
func RegisterRoutes(router *gin.Engine) {
	// Definimos el grupo de rutas relacionadas con usuarios
	// Esto se maneja por una función separada para mantener el código limpio y modular
	ConfigurarRutasUsuario(router)

	// Si necesitas rutas adicionales, las puedes agregar aquí
	// Ejemplo de ruta raíz
	router.GET("/", controllers.HomeHandler)

	// Ruta de prueba para verificar si el servidor está activo
	router.GET("/ping", controllers.PingHandler)

	// Puedes agregar más grupos de rutas y endpoints aquí, según sea necesario
}

// ConfigurarRutasUsuario configura las rutas relacionadas con usuarios
func ConfigurarRutasUsuario(router *gin.Engine) {
	usuarioGroup := router.Group("/users")
	{
		usuarioGroup.POST("/", controllers.CrearUsuario)         // Crear un nuevo usuario
		usuarioGroup.GET("/", controllers.ObtenerUsuarios)       // Obtener lista de usuarios
		usuarioGroup.GET("/:id", controllers.ObtenerUsuario)     // Obtener un solo usuario por ID
		usuarioGroup.PUT("/:id", controllers.ActualizarUsuario)  // Actualizar un usuario
		usuarioGroup.DELETE("/:id", controllers.EliminarUsuario) // Eliminar un usuario
	}
}
