# 🎮 4 in a Row (Connect Four) – Real-Time Multiplayer Game

This project is a real-time implementation of the classic **Connect Four (4 in a Row)** game built using **Go** for the backend and **React (Vite + Tailwind CSS)** for the frontend.

The game supports **real-time multiplayer gameplay using WebSockets**, with an intelligent **bot opponent** that automatically joins if no second player is found within a fixed timeout.

---

Live Demo :-  https://connect4-1-5gji.onrender.com/

## 🚀 Features

### 🧍 Player Matchmaking
- Players join the game by entering a username.
- If a second player joins within **10 seconds**, a **1v1 match** is started.
- If no second player joins, a **BOT** automatically starts the game with the waiting player.

---

### 🌐 Real-Time Gameplay
- Uses **WebSockets** for live, turn-based communication.
- Both players receive **instant board updates** after every move.
- Turn validation is handled on the backend to prevent invalid moves.

---

### 🤖 Competitive Bot
The bot is **not random** and follows basic strategy:
- Attempts to **win** if a winning move is available.
- **Blocks the opponent’s immediate winning move**.
- Prefers the **center column** when possible.
- Falls back to the first valid column if no strategic move exists.

---

### 🧠 Game Logic
- Standard **7 × 6 Connect Four board**.
- Supports:
  - Horizontal wins
  - Vertical wins
  - Diagonal wins (both directions)
- Draw detection when the board is full.
- Game state is maintained **in-memory** during active matches.

---

### 🏅 Leaderboard
- Tracks **number of wins per player** (including BOT).
- Leaderboard is maintained in memory.
- Exposed via a REST API endpoint:


## Author
Niranjan C B
