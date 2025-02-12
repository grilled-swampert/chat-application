import React, { Component } from 'react';
import './ChatHistory.scss';
import Message from '../Message/Message';

class ChatHistory extends Component {
  render() {
    // this.props.chatHistory is an array of messages.
    // We map over the array and create a new array of Message components.
    // Each message component is passed a message object.
    // The message object is the data of the message.
    // The key is the timestamp of the message.
    // The key is used by React to keep track of the Message components.
    // The key is important because it helps React to know which Message component to update.
    // The key should be unique for each Message component.
    // The key should not change over time.
    console.log(this.props.chatHistory);
    const messages = this.props.chatHistory.map(msg => <Message key={msg.timeStamp} message={msg.data} />); 

    return (
      <div className='ChatHistory'>
        <h2>Chat History</h2>
        {messages}
      </div>
    );
  };

}

export default ChatHistory;
