package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader: The Upgrader struct is used to upgrade an HTTP server connection to a WebSocket connection.
// ReadBufferSize: The size of the read buffer for the WebSocket connection.
// WriteBufferSize: The size of the write buffer for the WebSocket connection.
// Read Buffer: A buffer used to read data from the WebSocket connection.
// Write Buffer: A buffer used to write data to the WebSocket connection.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Upgrade: Upgrades the HTTP server connection to a WebSocket connection.
// w: The HTTP response writer.
// r: The HTTP request.
// Returns: A WebSocket connection and an error, if any.
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// CheckOrigin: A function that returns true to allow any origin to connect to the WebSocket server.
	// conn: The WebSocket connection.
	// err: An error, if any.
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil) // Upgrade the HTTP server connection to a WebSocket connection.
	if err != nil { // If an error occurs, log the error message and return nil.
		log.Println(err)
		return nil, err
	}

	return conn, nil // Return the WebSocket connection and nil.
}
