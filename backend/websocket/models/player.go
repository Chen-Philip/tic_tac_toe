package models

import (
	"encoding/json"
	"fmt"
	"log"

	"sync"

	"github.com/gorilla/websocket"
)

type Player struct {
	User_id   string
	Conn      *websocket.Conn
	Is_closed bool
	GameRoom  *GameRoom
	mu        sync.Mutex
}

func (p *Player) Read() {
	defer func() {
		p.GameRoom.Unregister <- p
	}()

	for {
		var msg Message
		err := p.Conn.ReadJSON(&msg)
		fmt.Println(msg)
		switch msg.Type {
		case MoveMessageType:
			var move MoveMessage
			json.Unmarshal(msg.Body, &move)
			if p.GameRoom.GameState.IsWin {
				p.Conn.WriteJSON(Message{Type: EndGameMessageType, Body: toRawJson("The game is over!")})
				continue
			}
			if p != p.GameRoom.PlayerTurn[p.GameRoom.GameState.Turn%2] {
				p.Conn.WriteJSON(Message{Type: TextMessageType, Body: toRawJson("It's not your turn!")})
				continue
			}
			if p.GameRoom.GameState.IsValidMove(move.X, move.Y) {
				p.GameRoom.GameState.MakeMove(move.X, move.Y)
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
