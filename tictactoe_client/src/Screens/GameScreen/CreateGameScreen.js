import React, { useState, useEffect, useRef } from "react";
import { Container, Header, Button, Form, Modal } from "semantic-ui-react";

function CreateGameScreen({ onSubmit }) {
  const [gameId, setgameId] = useState("");

  return (
    <Container textAlign="center" style={{ marginTop: "100px" }}>
      {/* Get the game room name */}
      <Header as="h1">
        Enter Game Room ID
      </Header>
      <Form>
        <Form.Input
          label="Game Room ID"
          placeholder="Enter room ID"
          value={gameId}
          onChange={(e) => setgameId(e.target.value)}
        />
      </Form>
      {/* Button for creating the game */}
      <Button
        primary
        onClick={() => onSubmit(gameId)}>
        Join / Create
      </Button>
    </Container>
  );
}

export default CreateGameScreen;