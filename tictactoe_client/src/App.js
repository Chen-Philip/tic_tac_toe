import React from "react";
import { Container, Header, Button } from "semantic-ui-react";
import "semantic-ui-css/semantic.min.css";
import AuthScreen from "./Screens/AuthScreen";
import MainScreen from "./Screens/MainScreen";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";


function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<AuthScreen />} />
        <Route path="/profile" element={<MainScreen />} />
      </Routes>
    </Router>
    // <Container textAlign="center" style={{ marginTop: "100px" }}>
    //   <AuthScreen/>
    // </Container> 
  );
}

export default App;
