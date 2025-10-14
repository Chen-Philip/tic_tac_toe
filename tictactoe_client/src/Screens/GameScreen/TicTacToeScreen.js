import React, { useState, useEffect } from "react";
import { Container, Header, Table } from "semantic-ui-react";
import GameAlertModal from "../../Components/GameAlertModel";
import { MessageType } from "../../Components/MessageType";

function TicTacToeScreen({ message, onCellClick, onExit }) {

  const [board, setBoard] = useState([
    [0, 0, 0],
    [0, 0, 0],
    [0, 0, 0],
  ]);
  const [modalMessage, setModalMessage] = useState("");
  const [modalButtonText, setModalButtonText] = useState("");
  const [showModal, setShowModal] = useState(false);
  const [messageType, setMessageType] = useState(-1);
  const [playerNum, setPlayerNum] = useState(-1);
  const [yourTurn, setYourTurn] = useState(false);
  const [hasOpponent, setHasOpponent] = useState(false);

  const handleModalClick = () => {
    if (messageType === MessageType.EndGameMessageType) {
      onExit();
    }
    setShowModal(false); // close modal
  };

  useEffect(() => {
    if (!message) return;
    
    let parsedMessage;
    try {
      parsedMessage = JSON.parse(message);
    } catch (err) {
      console.error("Failed to parse message:", err);
      return;
    }
    console.log(`type: ${parsedMessage.type}`)
    setMessageType(parsedMessage.type);
    switch (parsedMessage.type) {
      case MessageType.TextMessageType:
        setModalButtonText("Okay");
        setModalMessage(JSON.stringify(parsedMessage.body));
        setShowModal(true);
        break;
      case MessageType.EndGameMessageType:
        setModalButtonText("Leave Game");
        setModalMessage(JSON.stringify(parsedMessage.body));
        setShowModal(true);
        break;
      case MessageType.GameStateMessageType:
        setHasOpponent(true)
        setBoard(parsedMessage.body.board);
        setYourTurn(parsedMessage.body.Turn % 2 == playerNum);
        if (parsedMessage.body.IsWin) {
          let winText
          if (yourTurn) {
            winText = "You Win!"
          } else {
            winText = "You Lose!"
          }
          setModalButtonText("Leave Game");
          setMessageType(MessageType.EndGameMessageType);
          setModalMessage(winText);
          setShowModal(true);
        } 
        if (parsedMessage.body.Turn > 9) {
          setModalButtonText("Leave Game");
          setMessageType(MessageType.EndGameMessageType);
          setModalMessage("It's a tie!");
          setShowModal(true);
        }
        // setTurn(parsedMessage.body.Turn);
        break;
      case MessageType.PlayerTurnMessageType:
        setPlayerNum(parsedMessage.body.player);
        break;
      default:
        break;
    }
  }, [message]);

  if (!hasOpponent) {
    return (
      <Container textAlign="center" style={{ marginTop: "50px" }}>
        <Header as="h2">Waiting for Opponent...</Header>
      </Container>
    );
  }

  return (
    <Container textAlign="center" style={{ marginTop: "50px" }}>
      <Header as="h2">Tic Tac Toe</Header>
      {yourTurn ? (
        <Header as="h3">Your Turn</Header>
      ) : (
        <Header as="h3">Opponent's Turn</Header>
      )}
      {showModal && (
        <GameAlertModal
          msg={modalMessage}
          buttonText={modalButtonText}
          handleButtonClick={() => handleModalClick()}
        />
      )}

      <Table celled compact textAlign="center">
        <Table.Body>
          {board.map((row, rowIndex) => (
            <Table.Row key={rowIndex}>
              {row.map((cell, colIndex) => (
                <Table.Cell
                  key={colIndex}
                  style={{
                    fontSize: "2em",
                    width: "100px",
                    height: "100px",
                    borderLeft: colIndex === 0 ? "none" : "2px solid black", // vertical border
                    borderTop: rowIndex === 0 ? "none" : "2px solid black", // horizontal border
                  }}
                  onClick={() => onCellClick(rowIndex, colIndex)} // pass indices
                >
                  {cell === 0 ? "" : cell === 1 ? "X" : "O"}
                </Table.Cell>
              ))}
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
    </Container>
  );
}

export default TicTacToeScreen;
