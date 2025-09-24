package models

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Player struct {
	User_id  string
	Conn     *websocket.Conn
	GameRoom *GameRoom
	mu       sync.Mutex
}

func (p *Player) Read() {
	defer func() {
		p.GameRoom.Unregister <- p
		p.Conn.Close()
	}()

	for {
		msgType, body, err := p.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{
			Type: msgType,
			Body: string(body),
		}
		p.GameRoom.Broadcast <- message
		fmt.Println("Message recieved: ", message)
	}
}
