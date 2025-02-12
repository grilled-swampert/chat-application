import React, { Component } from 'react';
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';
import ChatInput from './components/ChatInput/ChatInput';
import './App.css';
import { connect, sendMsg } from './api';

// extends Component? What is this? 
// It is a way to create a class in JavaScript.
// React components are custom objects that inherit from React.Component.
class App extends Component {
  // The super keyword is used to access and call functions on an object's parent.
  // The constructor method is a special method for creating and initializing an object created with a class.
  // The constructor method is called automatically when a class is initiated, and it has to have the exact name "constructor", in fact, if you do not have a constructor method, JavaScript will add an invisible and empty constructor method.
  constructor(props) {
    super(props);
    this.state = { // The state object is where you store property values that belongs to the component.
      chatHistory: [] // chatHistory is an array that will store the chat messages.
    }
  }

  // componentDidMount() is invoked immediately after a component is mounted (inserted into the tree).
  // Initialization that requires DOM nodes should go here. If you need to load data from a remote endpoint, this is a good place to instantiate the network request.
  componentDidMount() { 
    connect((msg) => { // connect is a function that takes a callback function as an argument.
      console.log("New Message") // This will log "New Message" to the console.
      this.setState(prevState => ({ // The setState() method is used to update the state of the component. It will replace the old state with the new state.
        chatHistory: [...prevState.chatHistory, msg] // The chatHistory array will be updated with the new message.
      }))
      console.log(this.state); // This will log the state to the console.
    });
  }

  // send is a function that takes an event as an argument.
  send(event) {
    if (event.keyCode === 13) { // If the key pressed is the Enter key, then the following code will execute.
      sendMsg(event.target.value); // sendMsg is a function that takes the value of the input field as an argument.
      event.target.value = "";
    }
  }

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} />
      </div>
    );
  }
}

export default App;
