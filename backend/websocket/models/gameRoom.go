package models

import (
	"encoding/json"
	"fmt"
	"tictactoe/tic_tac_toe"
)

type GameRoom struct {
	Game_id    string
	GameState  tic_tac_toe.Game
	Register   chan *Player
	Unregister chan *Player // not sure if needed
	PlayerTurn []*Player
	Players    map[*Player]bool
	Broadcast  chan struct{}
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
			Broadcast:  make(chan struct{}),
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

				fmt.Println(len(gameRoom.Players))
				if len(gameRoom.Players) == 2 {
					fmt.Println("Start game")
					sendTurnMessage(gameRoom)
				}

				go newPlayer.Read()
			} else {
				newPlayer.Conn.WriteJSON(Message{
					Type: TextMessageType,
					Body: toRawJson("This game room is full!"),
				})
				newPlayer.Conn.Close()
			}
		case <-gameRoom.Unregister:
			fmt.Println("unregister")

			for p := range gameRoom.Players {
				delete(gameRoom.Players, p)
				fmt.Println("unregister 2")
				p.Conn.WriteJSON(Message{
					Type: TextMessageType,
					Body: toRawJson("Your opponent left. Game closed."),
				})
				p.Conn.Close()
			}

			fmt.Println("unregister 3")
			if len(gameRoom.Players) == 0 {
				fmt.Println("unregister 4")
				delete(rooms, gameRoom.Game_id)
				return
			}
			fmt.Println("unregister 5")
		case <-gameRoom.Broadcast:
			fmt.Println("Broadcast")
			message, _ := json.Marshal(GameStateMessage{
				Board: gameRoom.GameState.Board,
				IsWin: gameRoom.GameState.IsWin,
				Turn:  gameRoom.GameState.Turn,
			})
			fmt.Println("send Broadcast")
			broadcastMsg := Message{
				Type: GameStateMessageType,
				Body: message,
			}
			for p := range gameRoom.Players {
				if p.Conn.WriteJSON(broadcastMsg) != nil {
					fmt.Println("A broadcasting error has occurred")
				}
			}
		}
	}
}
