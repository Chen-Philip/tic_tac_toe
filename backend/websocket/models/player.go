package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Player struct {
	User_id   string
	Conn      *websocket.Conn
	Is_closed bool
	GameRoom  *GameRoom
}

func (p *Player) Read() {
	defer func() {
		// Game room closed, so disconnect player
		p.GameRoom.Unregister <- p
	}()

	for {
		// Reads the player's moves
		var msg Message
		err := p.Conn.ReadJSON(&msg)
		fmt.Println(msg)
		switch msg.Type {
		case MoveMessageType: // Player makes a move
			var move MoveMessage
			json.Unmarshal(msg.Body, &move)
			// Check if the game is over
			if p.GameRoom.GameState.IsWin {
				p.Conn.WriteJSON(Message{Type: EndGameMessageType, Body: toRawJson("The game is over!")})
				continue
			}
			// Check if the player is making a move on their turn
			if p != p.GameRoom.PlayerTurn[p.GameRoom.GameState.Turn%2] {
				p.Conn.WriteJSON(Message{Type: TextMessageType, Body: toRawJson("It's not your turn!")})
				continue
			}
			// Checks if the move is valid
			if p.GameRoom.GameState.IsValidMove(move.X, move.Y) {
				// Update the board
				p.GameRoom.GameState.MakeMove(move.X, move.Y)
				// Broadcast the move to the other player
				p.GameRoom.Broadcast <- struct{}{}
			} else {
				p.Conn.WriteJSON(Message{Type: TextMessageType, Body: toRawJson("That is not a valid move!")})
			}
		}

		if err != nil {
			log.Println(err)
			return
		}
	}
}
