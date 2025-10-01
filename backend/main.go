package main

import (
	"os"
	authRoutes "tictactoe/authentication/routes"
	wsRoutes "tictactoe/websocket/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Starts tic tac toe game in terminal:
	//s := tic_tac_toe.Game{}
	//s.StartTerminalGame()

	// Reads the system's environment variable called PORT
	port := os.Getenv("PORT")
	// Sets port to default value if it doesnt have a value
	if port == "" {
		port = "8000"
	}

	// gin.New() creates a new Gin engine without any middleware, unlike gin.default()
	// Gin engine: the main thing that runs the web server. It handles the router, which
	//		maps paths to handlers, middleware and configuration for running http server
	// Middleware: the middle steps/ code that runs before (sometimes after) the route handle
	// 			the middleman between the request and the handler function
	// router.Use(gin.Logger()): addes middleware gin.Logger to our engine
	// gin.Logger: jsut logs the requests
	router := gin.New()
	router.Use(gin.Logger())

	// Add middleware Cors so that frontend can talk to backend (need this for when frontend and backend are from different origins)
	// AllowHeaders: What headers the frontend sends that the server will accept
	//		Content-type: needed for sending json
	// 		Authorization: Needed for sending tokens, or else JWT and authorization
	// MaxAge: Kind of like a "remember me", gives time frame to not need the preflight request
	// Whenever a request is sent, a preflight request is sent to make sure the frontend can access the backend
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},                 // React frontend, what websties can bacvkend talk to
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"}, // Gives full method access
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{}, // Specify which headers can be read by frontend
		AllowCredentials: false,      // can the sites send cookies/auth info with requests, I dont need it because I use token apis
		MaxAge:           12 * time.Hour,
	}))

	authRoutes.AuthRoutes(router)
	authRoutes.UserRoutes(router)
	wsRoutes.GameRoomRoutes(router)

	// 1. router.GET("/api-1", ... registers a handler (function that processes incoming http
	//		requests and produces response) for HTTP get requests to the path api-1
	// 2. func(c *gin.Context) { ... } the inline handler function for the request
	// 3. gin.Context holds the request and respone info (header, params, JSON body ...
	// 4. c.JSON(200, gin.H{...} sends a JSON response back with status code 200
	// 5. gin.H{...} creates the json, but gin.H specifially maps a string to any dataype
	// Therefore the following code upon recieving a HTTP getrequest to the specified path will
	//		send a JSON response with 200 status code and a body of whatever gin.wH has
	// NOTE: THIS IS COMMENTED OUT SINCE API-1 AND API-2 IS NEVER USED
	//router.GET("/api-1", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"success": "Access granted for api-1"})
	//})
	//
	//router.GET("/api-2", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"success": "Access granted for api-2"})
	//})

	router.Run(":" + port)
}
