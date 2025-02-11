// Import the React library and Component base class
import { Component } from 'react';
import { connect, sendMsg } from '../api'; // Import WebSocket functions
import './App.css'; // Import CSS styles
import Header from './components/Header';
import ChatHistory from './components/ChatHistory/chatHistory' // Ensure ChatHistory component is imported

// Declare a class-based component by extending React's Component
class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: [],
    };
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message received:", msg);
      this.setState((prevState) => ({
        chatHistory: [...prevState.chatHistory, msg], // Proper state update
      }));
    });
  }

  send() {
    console.log("Sending message: hello"); // Log message before sending
    sendMsg("hello"); // Use the imported WebSocket function to send a message
  }

  // Render method defines the component's UI
  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        {/* Button with an event handler to call the `send` method */}
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}

// Export the component for use in other parts of the application
export default App;
