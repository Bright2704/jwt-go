package routes

import (
	controller "github.com/bright2704/jwt-api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/bright2704/jwt-api/middleware"
)

func UserRoutes (incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.AuthMiddleware())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}