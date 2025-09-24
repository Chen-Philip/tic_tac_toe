package models

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type GameRoom struct {
	Game_id    string
	Register   chan *Player
	Unregister chan *Player // not sure if needed
	Players    map[*Player]bool
	Broadcast  chan Message
}

func (gameRoom *GameRoom) StartGame() {
	for { // Runs an infinite loop
		select { // listens for channels to have data
		case newPlayer := <-gameRoom.Register: // Register channel has data
			if len(gameRoom.Players) < 2 {
				gameRoom.Players[newPlayer] = true
				fmt.Println("Player joined room ", gameRoom.Game_id)
			} else {
				newPlayer.Conn.WriteMessage(websocket.TextMessage, []byte("This game room is full!"))
				newPlayer.Conn.Close()
			}
		case player := <-gameRoom.Unregister:
			delete(gameRoom.Players, player)
			player.Conn.Close()
			fmt.Println("Player left room ", gameRoom.Game_id)

		case message := <-gameRoom.Broadcast:
			for player := range gameRoom.Players {
				if player.Conn.WriteMessage(message.Type, []byte(message.Body)) != nil {
					fmt.Println("A broadcasting error has occurred")
				}
			}
		}
	}
}
