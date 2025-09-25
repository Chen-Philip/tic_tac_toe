package models

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type GameRoom struct {
	Game_id    string
	Register   chan *Player
	Unregister chan *Player // not sure if needed
	PlayerTurn []*Player
	Players    map[*Player]bool
	Broadcast  chan Message
}

var rooms = make(map[string]*GameRoom)

func CreateOrGetGameRoom(gameId string) *GameRoom {
	if rooms[gameId] == nil {
		rooms[gameId] = &GameRoom{
			Game_id:    gameId,
			Register:   make(chan *Player),
			Unregister: make(chan *Player), // not sure if needed
			Players:    make(map[*Player]bool),
			PlayerTurn: make([]*Player, 0),
			Broadcast:  make(chan Message),
		}
	}

	go rooms[gameId].StartGame()

	return rooms[gameId]
}

func (gameRoom *GameRoom) StartGame() {
	for { // Runs an infinite loop
		select { // listens for channels to have data
		case newPlayer := <-gameRoom.Register: // Register channel has data
			if len(gameRoom.Players) < 2 {
				gameRoom.Players[newPlayer] = true
				gameRoom.PlayerTurn = append(gameRoom.PlayerTurn, newPlayer)
				fmt.Println("Player joined room ", len(gameRoom.Players))

				go newPlayer.Read()
			} else {
				newPlayer.Conn.WriteMessage(websocket.TextMessage, []byte("This game room is full!"))
				newPlayer.Conn.Close()
			}
		case player := <-gameRoom.Unregister:
			delete(gameRoom.Players, player)
			for p := range gameRoom.Players {
				p.Conn.WriteMessage(websocket.TextMessage, []byte("Your opponent left. Game closed."))
				p.Conn.Close()
			}

			fmt.Println("unregister")
			if len(gameRoom.Players) == 0 {
				delete(rooms, gameRoom.Game_id)
				return
			}
		case message := <-gameRoom.Broadcast:
			for p := range gameRoom.Players {
				if p.Conn.WriteMessage(message.Type, []byte(message.Body)) != nil {
					fmt.Println("A broadcasting error has occurred")
				}
			}
		}
	}
}
