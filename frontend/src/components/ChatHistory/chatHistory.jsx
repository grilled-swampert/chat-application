import React, { Component } from 'react';

class ChatHistory extends Component {
  render() {
    console.log(this.props.chatHistory);
    const messages = this.props.chatHistory.map((msg, index) => (
      <p key={index}>{msg?.data || JSON.stringify(msg)}</p> 
    ));

    return (
      <div className="bg-[#f7f7f7] w-full m-0 p-2 text-black">
        <h2 className="text-2xl text-center">Chat History</h2>
        <div className="space-y-2">{messages}</div>
      </div>
    );
  }
}

export default ChatHistory;