package main

import (
	"fmt"
	"net/http"

	"github.com/grilled-swampert/chat-application/backend/pkg/websocket"
)

// fmt: Provides formatted I/O functions (e.g., Println, Fprintf).
// net/http: Used to create an HTTP server and handle WebSocket connections.
// "github.com/grilled-swampert/chat-application/backend/pkg/websocket": Custom package that handles WebSocket connections.

// serveWs: Upgrades the HTTP server connection to a WebSocket connection.
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) { // The serveWs function takes a WebSocket pool, an HTTP response writer, and an HTTP request as arguments.
	fmt.Println("WebSocket Endpoint Hit") // Print a message to the console when a WebSocket connection is established.
	conn, err := websocket.Upgrade(w, r) // Upgrade the HTTP server connection to a WebSocket connection.
	if err != nil { 					// If an error occurs, print the error message to the console.
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{ // Create a new WebSocket client.
		Conn: conn, // Set the connection to the WebSocket connection.
		Pool: pool, // Set the pool to the WebSocket pool.
	}

	// What is a pool? A pool is a collection of WebSocket clients that are connected to the server.

	pool.Register <- client // Register the WebSocket client with the pool.
	client.Read() // Read messages from the WebSocket client.
}

// setupRoutes: Sets up the WebSocket route.
func setupRoutes() { 
	pool := websocket.NewPool() // Create a new WebSocket pool.
	go pool.Start() // Start the WebSocket pool.

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { // Handle WebSocket connections on the /ws route.
		serveWs(pool, w, r) // Upgrade the HTTP server connection to a WebSocket connection.
	})
}

func main() { // The main function is the entry point of the application.
	fmt.Println("Chat App") // Print a message to the console when the application starts.
	setupRoutes() // Set up the WebSocket route.
	http.ListenAndServe(":8080", nil) // Start the HTTP server on port 8080.
}
