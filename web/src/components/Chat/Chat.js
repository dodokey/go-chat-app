import React, { useState, useEffect, useRef } from "react";
import { useAuth } from "../../contexts/AuthContext";
import "./Chat.css";

function Chat() {
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState("");
  const { user } = useAuth();
  const messagesEndRef = useRef(null);
  const userUsername = user.username; // Extract username to avoid dependency warning

  useEffect(() => {
    // TODO: setup websocket, get history messages
    setMessages([
      { id: 1, text: "message 1", sender: "xxx" },
      { id: 2, text: "message 2", sender: userUsername },
      { id: 3, text: "message 3", sender: "xxx" },
      { id: 4, text: "message 4", sender: userUsername },
      { id: 5, text: "message 5", sender: "xxx" },
    ]);
  }, [userUsername]);

  useEffect(() => {
    scrollToBottom();
  }, [messages]);

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  };

  const handleSendMessage = (e) => {
    e.preventDefault();
    if (newMessage.trim()) {
      // TODO: setup websocket, send message to server
      setMessages([
        ...messages,
        { id: messages.length + 1, text: newMessage, sender: userUsername },
      ]);
      setNewMessage("");
    }
  };

  return (
    <div className="chat-container">
      <div className="messages-container">
        {messages.map((message) => (
          <div
            key={message.id}
            className={`message ${
              message.sender === userUsername ? "sent" : "received"
            }`}
          >
            <div className="avatar-container">
              <img
                src={message.sender === userUsername ? "2.jpeg" : "1.jpeg"}
                alt={message.sender}
                className="avatar"
              />
              <span className="username">{message.sender}</span>
            </div>
            <div className="message-content">{message.text}</div>
          </div>
        ))}
        <div ref={messagesEndRef} />
      </div>
      <form onSubmit={handleSendMessage} className="message-form">
        <input
          type="text"
          value={newMessage}
          onChange={(e) => setNewMessage(e.target.value)}
          placeholder="Type a message..."
          className="message-input"
        />
        <button type="submit" className="send-button">
          ➤
        </button>
      </form>
    </div>
  );
}

export default Chat;
