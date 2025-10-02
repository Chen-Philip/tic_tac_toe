package routes

import (
	controller "tictactoe/authentication/controllers"

	"github.com/gin-gonic/gin"
)

// A route maps HTTP methods and URL paths to a function handler
func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}
