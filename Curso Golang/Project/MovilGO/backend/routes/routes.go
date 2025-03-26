package routes

import (
	"movilGO/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/hello", controllers.GetHello)
	r.POST("/login", controllers.Login)
}
