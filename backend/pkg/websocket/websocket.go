package websocket

import (
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)


var jwtSecret = []byte("secret-key")

// Upgrader: The Upgrader struct is used to upgrade an HTTP server connection to a WebSocket connection.
// ReadBufferSize: The size of the read buffer for the WebSocket connection.
// WriteBufferSize: The size of the write buffer for the WebSocket connection.
// Read Buffer: A buffer used to read data from the WebSocket connection.
// Write Buffer: A buffer used to write data to the WebSocket connection.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true},
}

func verifyJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", err
	}
	return username, nil
}

// Upgrade: Upgrades the HTTP server connection to a WebSocket connection.
// w: The HTTP response writer.
// r: The HTTP request.
// Returns: A WebSocket connection and an error, if any.
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, string, error) {
	// CheckOrigin: A function that returns true to allow any origin to connect to the WebSocket server.
	// conn: The WebSocket connection.
	// err: An error, if any.
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "No Authorization header", http.StatusUnauthorized)
		return nil, "", nil
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	userId, err := verifyJWT(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return nil, "", err
	}

	conn, err := upgrader.Upgrade(w, r, nil) // Upgrade the HTTP server connection to a WebSocket connection.
	if err != nil { // If an error occurs, log the error message and return nil.
		log.Println(err)
		return nil, "", err
	}

	return conn, userId, nil // Return the WebSocket connection, userId and nil.
}
