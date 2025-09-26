package models

func sendTurnMessage(gameRoom *GameRoom) {
	turn := gameRoom.GameState.Turn

	gameRoom.PlayerTurn[turn%2].Conn.WriteJSON(Message{Type: TurnMessageType, Body: toRawJson("Your Turn!")})
	gameRoom.PlayerTurn[(turn+1)%2].Conn.WriteJSON(Message{Type: TurnMessageType, Body: toRawJson("Opponent's Turn!")})
}

func getGameStateMessage(gameRoom *GameRoom) {
	return
}
