package routes

import (
	controller "tictactoe/websocket/controllers"

	"github.com/gin-gonic/gin"
)

// For webscokets, we use a get rest to establish connection to a websocket. Afterwards, the connections is upgraded
// to a websocket (done by the handler), and future communications are sent over the websocket and not vie HTTP
func GameRoomRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/ws/:id", controller.GameControllerWsHandler()) // :[param] indicates a route parameter
}
