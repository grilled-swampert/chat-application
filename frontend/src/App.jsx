// Import the React library and Component base class
import React, { Component } from 'react';
import { connect, sendMsg } from '../api'; // Import WebSocket functions
import './App.css'; // Import CSS styles

// Declare a class-based component by extending React's Component
class App extends Component {
  // Constructor is called when the component is created
  constructor(props) {
    // Call the parent (Component) constructor to initialize the component
    super(props);

    // Initialize WebSocket connection when the component is created
    connect();

    // Bind the `send` method to ensure `this` refers to the class instance
    this.send = this.send.bind(this);
  }

  // Custom method to send a WebSocket message
  send() {
    console.log("Sending message: hello"); // Log message before sending
    sendMsg("hello"); // Use the imported WebSocket function to send a message
  }

  // Render method defines the component's UI
  render() {
    return (
      <div className='App'>
        {/* Button with an event handler to call the `send` method */}
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}

// Export the component for use in other parts of the application
export default App;
