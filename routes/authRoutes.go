package routes

import (
	"shipment/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.GET("/users", controllers.GetUsers)
	}
}
