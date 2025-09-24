package helpers

import "tictactoe/websocket/models"

var rooms = make(map[string]*models.GameRoom)

func CreateOrGetGameRoom(gameId string) *models.GameRoom {
	if rooms[gameId] == nil {
		rooms[gameId] = &models.GameRoom{
			Game_id:    gameId,
			Register:   make(chan *models.Player),
			Unregister: make(chan *models.Player), // not sure if needed
			Players:    make(map[*models.Player]bool),
			Broadcast:  make(chan models.Message),
		}
	}

	go rooms[gameId].StartGame()

	return rooms[gameId]
}
