🎮 4 in a Row (Connect Four) – Real-Time Multiplayer Game

A real-time implementation of the classic 4 in a Row (Connect Four) game built using Go (Golang) for the backend and React (Vite) for the frontend.
The game supports player vs player and player vs bot gameplay using WebSockets.

🚀 Tech Stack
Backend

Go (Golang)

Gorilla WebSocket

In-memory game state

HTTP server for REST endpoints

Frontend

React

Vite

WebSocket client

Minimal UI (functional focus)

🕹️ Game Features (Implemented)
✅ Real-Time Gameplay

Uses WebSockets for real-time, turn-based play

Board updates are broadcast instantly after every move

Supports 7×6 grid with gravity-based disc drops

✅ Player Matchmaking

Player joins using a username

If another player joins within 10 seconds, a 1v1 game starts

If no player joins within 10 seconds, a BOT automatically starts the game

🤖 Competitive Bot

The bot is not random and follows basic strategic logic:

Tries to win immediately if possible

Blocks the opponent’s immediate winning move

Prefers center column

Falls back to the first valid column

🏆 Win Detection

Detects wins in:

Horizontal

Vertical

Diagonal (both directions)

Sends final board state before announcing the winner

📊 Leaderboard (In-Memory)

Tracks number of wins per player

Includes both human players and the BOT

Exposed via REST API:

GET /leaderboard


Displayed on the frontend UI after each game

⚠️ Leaderboard data is not persistent and resets on server restart.

🔁 Restart Game

Players can restart and rejoin a new game after completion

🖥️ Frontend UI

Enter username

Join game

Click columns to drop discs

See:

Live board updates

Turn indicator

Game result

Leaderboard

Styling is minimal; focus is on functionality.

🗂️ Project Structure
Connect4/
├── frontend/
│   ├── src/
│   │   └── App.jsx
│   └── dist/           # Production build (served by Go)
│
├── backend/
│   ├── main.go
│   ├── game.go
│   ├── bot.go
│   ├── matchmaking.go
│   ├── leaderboard.go
│   └── types.go
│
└── README.md

▶️ How to Run Locally
1️⃣ Build Frontend
cd frontend
npm install
npm run build


This generates the dist/ folder.

2️⃣ Run Backend Server
cd backend
go build -tags netgo -ldflags "-s -w" -o app
./app


Server runs on:

http://localhost:8080

3️⃣ Open the App

Open your browser and go to:

http://localhost:8080

📌 API Endpoints
WebSocket
GET /ws?username=<your_name>

Leaderboard
GET /leaderboard


Returns:

[
  {
    "username": "Player1",
    "wins": 2
  }
]

🎯 Why Go?

Strong concurrency model (goroutines)

Excellent WebSocket support

High performance for real-time systems

Simple and clean backend architecture

📌 Author
Niranjan C B
