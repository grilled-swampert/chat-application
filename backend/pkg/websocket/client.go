package websocket

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Client struct: Represents a WebSocket client.
// ID: A unique identifier for the client.
// Conn: The WebSocket connection.
// Pool: The WebSocket pool that the client belongs to.
// mu: A mutex to synchronize access to the client.
// What is a mutex? A mutex is a synchronization primitive that is used to protect shared resources from concurrent access.
// Synchonization primitive? A synchronization primitive is a mechanism that is used to coordinate access to shared resources in a concurrent system.
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex
}

// Message struct: Represents a WebSocket message.
// Type: The type of message (e.g., text, binary).
// Body: The content of the message.
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

// Read: Reads messages from the WebSocket client.
// What is a WebSocket client? A WebSocket client is a program that establishes a connection to a WebSocket server and sends and receives messages.
// What is a WebSocket server? A WebSocket server is a program that listens for incoming WebSocket connections and handles messages sent by clients.
func (c *Client) Read() {
	defer func() { // Defer the execution of the following code until the function returns.
		c.Pool.Unregister <- c // Unregister the client from the pool when the function returns.
		c.Conn.Close() 	   // Close the WebSocket connection when the function returns.
	}()

	// Infinite loop to read messages from the WebSocket client.
	for {
		messageType, p, err := c.Conn.ReadMessage() // Read a message from the WebSocket connection.
		if err != nil { // If an error occurs, print the error message to the console and return from the function.
			log.Println(err) 
			return
		}
		
		// Create a new message with the message type and body.
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message // Broadcast the message to all clients in the pool.
		fmt.Printf("Message Received: %+v\n", message) // Print the received message to the console.
	}
}
