package models

import "encoding/json"

const (
	TextMessageType       = 0
	MoveMessageType       = 1
	GameStateMessageType  = 2
	EndGameMessageType    = 3
	PlayerTurnMessageType = 4
)

// JSON for a message
type Message struct {
	Type int             `json:"type"`
	Body json.RawMessage `json:"body"`
}

// JSON to describe the move
type MoveMessage struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// JSON for the game state
type GameStateMessage struct {
	Board [3][3]int `json:"board"`
	IsWin bool      `json:"IsWin"`
	Turn  int       `json:"Turn"`
}

// JSON for telling the client what player they are
type PlayerTurnMessage struct {
	Player int `json:"player"`
}

// Converts a struct to a json
func toRawJson(v interface{}) json.RawMessage {
	body, _ := json.Marshal(v)
	return body
}
