package models

import (
	"encoding/json"
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
			Unregister: make(chan *Player),
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
				// Tells the client which player they are
				message, _ := json.Marshal(PlayerTurnMessage{
					Player: len(gameRoom.Players),
				})
				newPlayer.Conn.WriteJSON(Message{
					Type: PlayerTurnMessageType,
					Body: message,
				})
				// Add player to the room
				gameRoom.Players[newPlayer] = true
				gameRoom.PlayerTurn = append(gameRoom.PlayerTurn, newPlayer)
				// Create goroutine to read player moves
				go newPlayer.Read()
				// Let the players know the game started
				if len(gameRoom.Players) == 2 {
					gameRoom.Broadcast <- struct{}{}
				}
			} else {
				// The room is already full
				newPlayer.Conn.WriteJSON(Message{
					Type: EndGameMessageType,
					Body: toRawJson("This game room is full!"),
				})
				newPlayer.Conn.Close()
			}
		case <-gameRoom.Unregister: // Player leaves the game
			// Close the game room
			for p := range gameRoom.Players {
				// Removes player from the gameroom's player list
				delete(gameRoom.Players, p)
				// Disconnects the other player
				p.Conn.WriteJSON(Message{
					Type: EndGameMessageType,
					Body: toRawJson("Your opponent left. Game closed."),
				})
				p.Conn.Close()
			}
			// Removes the room from the list of gamerooms
			if len(gameRoom.Players) == 0 {
				delete(rooms, gameRoom.Game_id)
				return
			}
		case <-gameRoom.Broadcast: // Notify all players of the move
			// Creates message
			message, _ := json.Marshal(GameStateMessage{
				Board: gameRoom.GameState.Board,
				IsWin: gameRoom.GameState.IsWin,
				Turn:  gameRoom.GameState.Turn,
			})
			broadcastMsg := Message{
				Type: GameStateMessageType,
				Body: message,
			}
			// Sends the message
			for p := range gameRoom.Players {
				p.Conn.WriteJSON(broadcastMsg)
			}
		}
	}
}
