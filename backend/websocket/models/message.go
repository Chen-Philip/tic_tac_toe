package models

import "encoding/json"

const (
	TextMessageType      = 0
	MoveMessageType      = 1
	GameStateMessageType = 2
	EndGameMessageType	 = 3
)

type Message struct {
	Type int             `json:"type"`
	Body json.RawMessage `json:"body"`
}

type MoveMessage struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type GameStateMessage struct {
	Board [3][3]int `json:"board"`
	IsWin bool      `json:"IsWin"`
	Turn  int       `json:"Turn"`
}

func toRawJson(v interface{}) json.RawMessage {
	body, _ := json.Marshal(v)
	return body
}
