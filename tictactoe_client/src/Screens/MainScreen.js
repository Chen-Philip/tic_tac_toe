import React from "react";
import { Container, Header, Button } from "semantic-ui-react";
import { useLocation, useNavigate } from "react-router-dom";

function MainScreen() {
	const location = useLocation();
	const navigate = useNavigate();
  const user = location.state;
  console.log(user)
  
    return (
      <Container textAlign="center" style={{ marginTop: "100px" }}>
				
				{user ? (
					<Header as="h1">
						Welcome, {user.first_name} {user.last_name}
					</Header>
				) : (
					<Header>No user info</Header>
				)}
        <Button
					primary
					style={{ marginTop: "20px" }}
					onClick={() => navigate(-1)} // goes back to previous page
				>
					Back
				</Button>
      </Container>
    );
}

export default MainScreen;