import React from "react";
import { Container, Header, Button } from "semantic-ui-react";
import "semantic-ui-css/semantic.min.css";
import AuthScreen from "./AuthScreen";

function App() {
  return (
    <Container textAlign="center" style={{ marginTop: "100px" }}>
      <AuthScreen/>
    </Container> 
  );
}

export default App;
