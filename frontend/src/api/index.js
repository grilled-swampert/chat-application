// api/index.js
var socket = new WebSocket('ws://localhost:8080/ws');

// connect is a function that takes a callback function as an argument.
let connect = (cb) => {
  console.log("connecting")

  // onopen is an event handler that is called when the WebSocket connection's readyState changes to OPEN.
  socket.onopen = () => {
    console.log("Successfully Connected");
  }

  // onmessage is an event handler that is called when a message is received from the server.
  socket.onmessage = (msg) => {
    console.log("Message from WebSocket: ", msg);
    cb(msg); // cb is a callback function that takes msg as an argument.
  }

  // onclose is an event handler that is called when the WebSocket connection's readyState changes to CLOSED.
  socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event)
  }

  // onerror is an event handler that is called when an error occurs on the WebSocket.
  socket.onerror = (error) => {
    console.log("Socket Error: ", error)
  }
};

// sendMsg is a function that takes msg as an argument.
let sendMsg = (msg) => {
  console.log("sending msg: ", msg);
  socket.send(msg); // The send() method is used to send data to the server.
};

export { connect, sendMsg };
