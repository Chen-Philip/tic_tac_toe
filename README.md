Project Description:

This project is a simple full-stack project for an online tic-tac-toe game. The backend is build with GO, using the Gin framework and Gorilla WebSocket. The frontend is built using React, along with the Semantic UI React library.
NOTE: Front-end is developed using ChatGPT, since it wasn't the main focus of this project

Goal:

The main goal / focus of this project is to learn backend and familliarize myself with Go. 

Steps:
For this project, I broke it down into 3 main stages:
1. Developing the tic-tac-toe game logic
2. Developing the backend
3. Developing the frontend

Tic-Tac-Toe Game logic:
The primary goal for this section was to familiarize myself with Go, such as its syntax. I first made the game so that it can be played in the terminal,
where the user would input the moves as x and y coordinates. The program will then display the resulting board. This part of the project is also responsible
for the game logic and game state, such as validing moves and managing player turns, and can be used later on by the backend.

Backend:
The backend consists of 2 main sections:
1. Authenticaiton
   This project is my first introduction to backend, so I had to self-learn a lot of the fundamentals. Authentication was the first backend feature I implemented for this project and at this point,
   I had 0 experience with backend development. So, for this feature, I closely follwed [youtube tutorial](https://www.youtube.com/watch?v=Cr3BiwGN2Tg) by Akhill. As a result, most, if not all of the code, is directly from the youtube tutorial.
   However, I didn't just blindly copy what he did, but rather I took time to understand each line of code did andthe though process behind it. In fact, throughout the codebase, you can see the comments where I explained what each part of the code
   does to ensure my understanding. Thanks to that, I was able to learn a lot about the basics of backend development and how to use the GIN framework.
3. Game Creation /  Real-time game logic
   After creating the authentication, I'm more familliar with backend developement, so I didn't need to rely as heavily on youtube tutorials. The main challenge in this stage was implementing real-time gameplay, which required using WebSockets
   instead of traditional API calls, which I had no experience in. So I used this [Websocket Youtube Tutorial](https://www.youtube.com/watch?v=_hFPoXoMwXQ) by Akhill as reference for my backend logic. By learning "Pool" strucutre used in the
   video and modifying it to suit my needs, along with my exepreience I gained from implmeenting authenticaiton, I was able to build the real-time gameplay logic without relying to heaviliy on the youtube tutorial.
   This tutorial also taught me how to connect my backend to my frontend later on.

Frontend
Since front-end wasn't the main
