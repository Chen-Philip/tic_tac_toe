import React from "react";
import { Container, Header, Button } from "semantic-ui-react";
import "semantic-ui-css/semantic.min.css";
import AuthScreen from "./Screens/AuthScreen";
import MainScreen from "./Screens/MainScreen";
import GameScreen from "./Screens/GameScreen/GameScreen"
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";


function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<AuthScreen />} />
        <Route path="/main" element={<MainScreen />} />
        <Route path="/tictactoe" element={<GameScreen />} />
      </Routes>
    </Router>
  );
}

export default App;
