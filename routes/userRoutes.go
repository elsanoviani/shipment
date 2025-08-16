package routes

import (
	"shipment/controllers"
	"shipment/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("/register", controllers.RegisterUser)
		users.POST("/login", controllers.LoginUser)
		users.GET("/", middlewares.AuthMiddleware(), controllers.GetUsers)
		users.PUT("/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
		users.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)
	}
}
