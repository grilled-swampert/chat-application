package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024, // Setting the size of the read buffer.
  WriteBufferSize: 1024, // Setting the size of the write buffer.
  CheckOrigin: func(r *http.Request) bool { return true }, 
}

// A reader which will listen for messages from the client.
func reader(conn *websocket.Conn) {
  for {
    messageType, p, err := conn.ReadMessage() // Reading the message from the client.
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(string(p)) // Printing the message.
    if err := conn.WriteMessage(messageType, p); err != nil { // Writing the message back to the client.
      fmt.Println(err)
      return
    }
  }
}

// Defining the function that will be called when the endpoint is hit.
func serverSetup(w http.ResponseWriter, r *http.Request) {
  fmt.Printf(r.Host) // Printing the host of the request.

  // Upgrading the connection to a WebSocket connection.
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Listening indefinitely for messages from the client.
  reader(conn)
}

func setupRoutes() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Simple Server")
  })
  http.HandleFunc("/ws", serverSetup) // Defining a new endpoint whcih maps to the function.
}

func main() {
  setupRoutes()
  http.ListenAndServe(":8080", nil)
}