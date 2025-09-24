package routes

import (
	controller "tictactoe/authentication/controllers"
	"tictactoe/authentication/middleware"

	"github.com/gin-gonic/gin"
)

// UserRoutes We are using middleware here to ensure that both routes here are protected. Authroutes aren't protected
// because the user does't have a token yet. After loging, theyll have a token that they should use.
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
