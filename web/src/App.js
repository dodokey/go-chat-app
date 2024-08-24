import React, { useState } from "react";
import "./App.css";

function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [loginStatus, setLoginStatus] = useState(null);


  const handleSubmit = async (event) => {
    event.preventDefault();
    console.log("Username:", username);
    console.log("Password:", password);
    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
      });

      if (response.ok) {
        const data = await response.json();
        setLoginStatus(`Welcome, ${data.username}!`);
      } else {
        setLoginStatus('Login failed: Invalid username or password');
      }
    } catch (error) {
      setLoginStatus('Login failed: ' + error.message);
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>Login</h1>
        <form onSubmit={handleSubmit}>
          <div>
            <label>
              Username:
              <input
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
              />
            </label>
          </div>
          <div>
            <label>
              Password:
              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </label>
          </div>
          <div>
            <button type="submit">Login</button>
          </div>
        </form>
      </header>
    </div>
  );
}

export default App;
