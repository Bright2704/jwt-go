package routes

import (
	controller "github.com/bright2704/jwt-api/controllers"
	"github.com/gin-gonic/gin"
)


func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}