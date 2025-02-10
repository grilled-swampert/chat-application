// Create a new WebSocket instance and connect to the server at the specified endpoint
var socket = new WebSocket('ws://localhost:8080/ws');

let connect = () => {
    // Log an attempt to connect to the WebSocket server
    console.log("Attempting Connection...");
    
    // Event triggered when the connection is successfully established
    socket.onopen = () => {
        console.log("Successfully Connected");
    };
    
    // Event triggered when a message is received from the server
    socket.onmessage = msg => {
        // Log the received message
        console.log(msg);
    }

    // Event triggered when the connection is closed
    socket.onclose = event => {
        // Log the reason for connection closure
        console.log("Socket Closed Connection: ", event);
    }

    // Event triggered when an error occurs in the WebSocket connection
    socket.onerror = error => {
        // Log the error details
        console.log("Socket Error: ", error);
    }
}

// Function to send a message to the WebSocket server
let sendMsg = msg => {
    // Log the message being sent
    console.log("sending msg: ", msg);
    // Send the message through the WebSocket connection
    socket.send(msg);
}

// Export the connect and sendMsg functions so they can be used in other modules
export { connect, sendMsg };
