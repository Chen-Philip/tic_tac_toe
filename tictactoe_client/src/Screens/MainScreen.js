import React, { useState, useEffect, useRef } from "react";
import { Container, Header, Button, Form, Modal } from "semantic-ui-react";
import { useLocation, useNavigate } from "react-router-dom";

function MainScreen() {
  const location = useLocation();
  const navigate = useNavigate();
  const user = location.state;
  return (
    <Container textAlign="center" style={{ marginTop: "100px" }}>
      {/* Welcome user text */}
      {user ? (
        <Header as="h1">
          Welcome, {user.first_name} {user.last_name}
        </Header>
      ) : (
        <Header>No user info</Header>
      )}

      {/* Back Button */}
      <Button
        primary
        style={{ marginTop: "20px" }}
        onClick={() => navigate(-1)} // goes back to previous page
      >
        Back
      </Button>

      {/* Play Button */}
      <Button
        secondary
        style={{ marginTop: "20px", marginLeft: "10px" }}
        onClick={() => navigate("/tictactoe")}
      >
        Play
      </Button>
    </Container>
  );
}

export default MainScreen;