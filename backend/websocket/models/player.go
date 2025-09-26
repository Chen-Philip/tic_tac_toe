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
			fmt.Println("reading MoveMessage")
			var move MoveMessage
			json.Unmarshal(msg.Body, &move)
			fmt.Println(p.GameRoom.GameState.Turn)
			fmt.Println(move)
			if p.GameRoom.GameState.IsWin {
				p.Conn.WriteJSON(Message{Type: TurnMessageType, Body: toRawJson("The game is over!")})
				continue
			}
			if p != p.GameRoom.PlayerTurn[p.GameRoom.GameState.Turn%2] {
				fmt.Println("wrong turn")
				p.Conn.WriteJSON(Message{Type: TurnMessageType, Body: toRawJson("It's not your turn!")})
				continue
			}
			fmt.Println("right turn")
			if p.GameRoom.GameState.IsValidMove(move.X, move.Y) {
				fmt.Println("valid move")
				p.GameRoom.GameState.MakeMove(move.X, move.Y)
				p.GameRoom.Broadcast <- struct{}{}
			} else {
				fmt.Println("not valid")
				p.Conn.WriteJSON(Message{Type: TurnMessageType, Body: toRawJson("That is not a valid move!")})
			}
		}

		if err != nil {
			fmt.Println("error")
			log.Println(err)
			return
		}
	}
}
