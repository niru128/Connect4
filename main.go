package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWS(w http.ResponseWriter, r *http.Request) {

	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	player := &Player{
		ID:       uuid.New().String(),
		Username: username,
		Conn:     conn,
	}

	game := MatchPlayer(player)

	for game == nil {
		conn.WriteJSON(map[string]string{
			"type": "waiting",
		})
		time.Sleep(500 * time.Millisecond)
		game = playerGames[player.Username]
	}

	for {
		var msg WSMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			return
		}

		if game == nil {
			game = playerGames[player.Username]
			if game == nil {
				continue
			}
		}

		if msg.Type == "move" && game.Turn == player.Symbol {

			_, ok := game.DropDisc(msg.Column)
			if !ok {
				continue
			}

			// Human win check
			if game.CheckWin(player.Symbol) {

				RecordWin(player.Username)

				for _, p := range game.Players {
					if p.Conn != nil {
						p.Conn.WriteJSON(map[string]interface{}{
							"type":  "state",
							"board": game.Board,
							"turn":  game.Turn,
						})
					}
				}

				for _, p := range game.Players {
					if p.Conn != nil {
						p.Conn.WriteJSON(map[string]interface{}{
							"type":   "win",
							"winner": player.Username,
						})
					}
				}

				return
			}

			// Switch turn after human move
			game.SwitchTurn()

			for _, p := range game.Players {
				if p.Conn != nil {
					p.Conn.WriteJSON(map[string]interface{}{
						"type":  "state",
						"board": game.Board,
						"turn":  game.Turn,
					})
				}
			}

			nextPlayer := game.Players[game.Turn]
			if nextPlayer != nil && nextPlayer.Username == "BOT" {

				col := BotMove(game, game.Turn)
				if col == -1 {
					return
				}

				game.DropDisc(col)

				// Bot win check
				botSymbol := 2 // BOT is always player 2

				if game.CheckWin(botSymbol) {

					RecordWin("BOT")
					for _, p := range game.Players {
						if p.Conn != nil {
							p.Conn.WriteJSON(map[string]interface{}{
								"type":  "state",
								"board": game.Board,
								"turn":  game.Turn,
							})
						}
					}

					for _, p := range game.Players {
						if p.Conn != nil {
							p.Conn.WriteJSON(map[string]interface{}{
								"type":   "win",
								"winner": "BOT",
							})
						}
					}

					return
				}

				game.SwitchTurn()

				for _, p := range game.Players {
					if p.Conn != nil {
						p.Conn.WriteJSON(map[string]interface{}{
							"type":  "state",
							"board": game.Board,
							"turn":  game.Turn,
						})
					}
				}

			}
		}
	}
}

func main() {

	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", HandleWS)
	http.HandleFunc("/leaderboard", HandleLeaderboard)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
