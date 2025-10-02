import React, { useState, useRef } from "react";
import { Container, Header, Button, Form, Modal } from "semantic-ui-react";
import { useLocation, useNavigate } from "react-router-dom";
import CreateGameScreen from "./CreateGameScreen";
import TicTacToeScreen from "./TicTacToeScreen";
import { MessageType } from "../../Components/MessageType";

function GameScreen() {
	const location = useLocation();
	const navigate = useNavigate();
  // const user = location.state;

  const [showGameCreationScreen, setShowGameCreationScreen] = useState(true);
  // const [showGameScreen, setShowGameScreen] = useState(false);
  const [wsConnected, setWsConnected] = useState(false);
  const [messages, setMessages] = useState("");
  const ws = useRef(null); // store websocket instance

  const WS_ENDPOINT = "ws://localhost:9000/ws";

  const handleBackClick = () => {
    if (ws.current) {
      ws.current.close();
    }
    navigate(-1);
  };

  const handleJoinGame = (gameId) => {
    // Here you would normally connect to your game backend
    console.log("Joining/Creating room:", gameId);
    console.log(`${WS_ENDPOINT}/${gameId}`);
    ws.current = new WebSocket(`${WS_ENDPOINT}/${gameId}`);

    ws.current.onopen = () => {
      console.log("WebSocket connected");
      setWsConnected(true);
      setShowGameCreationScreen(false); // close modal
    };

    ws.current.onclose = () => {
      console.log("WebSocket disconnected");
      setWsConnected(false);
    };

    ws.current.onmessage = (event) => {
      console.log("Message from server:", event.data);
      setMessages(event.data);
    };
    
  };

  const handleCellClick = (row, col) => {
    console.log("Clicked cell:", row, col);
    
    if (ws.current && ws.current.readyState === WebSocket.OPEN) {
      const message = {
        type: MessageType.MoveMessageType, // MoveMessageType
        body: { x: row, y: col },
      };
      ws.current.send(JSON.stringify(message));
    }
  };

  
    return (
      <Container textAlign="center" style={{ marginTop: "100px" }}>
				
				{showGameCreationScreen ? (
					<CreateGameScreen
            onSubmit={handleJoinGame}
          />
				) : (
					<TicTacToeScreen message={messages} onCellClick={handleCellClick} onExit={handleBackClick} />
				)}

        <Button
					primary
					style={{ marginTop: "20px" }}
					onClick={ handleBackClick } // goes back to previous page
				>
					Back
				</Button>
      </Container>
    );
}

export default GameScreen;