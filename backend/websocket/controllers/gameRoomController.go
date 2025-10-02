package controllers

import (
	"fmt"
	"log"
	"net/http"
	"tictactoe/websocket/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(c *gin.Context) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}

func GameControllerWsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		gameID := c.Param("id")
		fmt.Println("Gameroom ", gameID, ": websocket endpoint reached")

		gameRoom := models.CreateOrGetGameRoom(gameID)

		conn, err := Upgrade(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		player := &models.Player{
			User_id:   "temptemp",
			Is_closed: false,
			Conn:      conn,
			GameRoom:  gameRoom,
		}

		gameRoom.Register <- player
	}
}
