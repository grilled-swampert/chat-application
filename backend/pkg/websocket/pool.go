package websocket

import "fmt"

// Pool struct: Represents a WebSocket connection pool.
// Register: A channel to register new WebSocket clients.
// Unregister: A channel to unregister WebSocket clients.
// Clients: A map of WebSocket clients.
// Broadcast: A channel to broadcast messages to all clients in the pool.
// chan *Client: A channel that sends and receives pointers to WebSocket clients.
// what is chan? A channel is a communication mechanism that allows goroutines to send and receive values.
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

// NewPool: Creates a new WebSocket connection pool.
// Returns a pointer to the WebSocket connection pool.
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

// Start: Starts the WebSocket connection pool.
func (pool *Pool) Start() {
	for { // Infinite loop to handle WebSocket clients.
		select { // Select statement to handle multiple channels.
		case client := <-pool.Register: // If a client is registered with the pool, add the client to the pool.
			pool.Clients[client] = true // Add the client to the pool.
			fmt.Println("Size of Connection Pool: ", len(pool.Clients)) // Print the size of the connection pool to the console.
			for client, _ := range pool.Clients { // Iterate over all clients in the pool.
				fmt.Println(client) // Print the client to the console.
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."}) // Send a message to the client.
			} 
			break // Exit the loop.
		case client := <-pool.Unregister: // If a client is unregistered from the pool, remove the client from the pool.
			delete(pool.Clients, client) // Remove the client from the pool.
			fmt.Println("Size of Connection Pool: ", len(pool.Clients)) // Print the size of the connection pool to the console.
			for client, _ := range pool.Clients { // Iterate over all clients in the pool.	
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."}) // Send a message to the client.
			}
			break
		case message := <-pool.Broadcast: // If a message is broadcast to the pool, send the message to all clients in the pool.
			fmt.Println("Sending message to all clients in Pool") // Print a message to the console.
			for client, _ := range pool.Clients { // Iterate over all clients in the pool.
				if err := client.Conn.WriteJSON(message); err != nil { // If an error occurs while sending the message, print the error message to the console.
					fmt.Println(err)
					return
				}
			}
		}
	}
}
