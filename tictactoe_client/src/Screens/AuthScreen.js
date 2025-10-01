import React, { useState } from "react";
import axios from "axios";
import { Container, Form, Button, Header } from "semantic-ui-react";
import { useNavigate } from "react-router-dom";

const endpoint = "http://localhost:9000";

function AuthScreen() {
  // for login
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  // for sign-up
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  // Shared states
  const [isLogin, setIsLogin] = useState(true);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault(); // Prevent browser's default behavior
    
    let authData = {}
    let authPath = ""

    if (isLogin) {
      authData = {
        "Username": username,
        "Password": password,
      }
      authPath = "/users/login"
    } else {
      // Check if passwords match
      if (password !== confirmPassword) {
        alert("Passwords do not match!");
        return;
      }

      authData = {
        "First_name": firstName,
        "Last_name": lastName,
        "Username": username,
        "Password": password,
      }
      authPath = "/users/signup"
    }

    // Send the response
    try {
      const response = await axios.post(`${endpoint}${authPath}`, authData);
      console.log(" Login successful:", response.data);
      navigate("/main", { state: response.data });
    } catch (error) {
      console.error("Login failed:", error.response?.data || error.message);
      alert("Login/Signup failed. Please try again.");
    }
  };

  return (
    <Container textAlign="center" style={{ marginTop: "100px" }}>
      <Header as="h2">Login</Header>
      <Form onSubmit={handleSubmit}>
        {!isLogin && (
          <>
            <Form.Input
              icon="user"
              iconPosition="left"
              label="First Name"
              placeholder="Enter your first name"
              type="text"
              value={firstName}
              onChange={(e) => setFirstName(e.target.value)}
              required
            />

            <Form.Input
              icon="user"
              iconPosition="left"
              label="Last Name"
              placeholder="Enter your last name"
              type="text"
              value={lastName}
              onChange={(e) => setLastName(e.target.value)}
              required
            />
          </>
        )}

        <Form.Input
          label="Username"
          type="text"
          placeholder="Enter your username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />
        <Form.Input
          label="Password"
          type="password"
          placeholder="Enter your password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />

        {!isLogin && (
            <Form.Input
              icon="lock"
              iconPosition="left"
              label="Confirm Password"
              placeholder="Re-enter your password"
              type="password"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
              required
            />
          )
        }

        <Button primary type="submit">
          {isLogin ? "Login" : "Sign Up"}
        </Button>

        <Button
          basic
          color="blue"
          fluid
          style={{ marginTop: "10px" }}
          onClick={() => setIsLogin(!isLogin)}
        >
          {isLogin
            ? "Don't have an account? Sign Up"
            : "Already have an account? Login"}
        </Button>
      </Form>
    </Container>
  );
}

export default AuthScreen;
