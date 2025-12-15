package main

import (
	"sync"
	"time"
)

var waitingPlayer *Player
var mutex sync.Mutex
var playerGames = make(map[string]*Game)

func MatchPlayer(p *Player) *Game {
	mutex.Lock()
	defer mutex.Unlock()

	if waitingPlayer == nil {
		waitingPlayer = p
		go StartBotAfterTimeOut(p)
		return nil
	}

	// Create game
	game := &Game{
		Turn:  1,
		Board: [ROWS][COLS]int{},
		Players: map[int]*Player{
			1: waitingPlayer,
			2: p,
		},
	}

	waitingPlayer.Symbol = 1
	p.Symbol = 2

	// 🔗 Register BOTH players
	playerGames[waitingPlayer.Username] = game
	playerGames[p.Username] = game

	// 🔔 Notify BOTH players
	waitingPlayer.Conn.WriteJSON(map[string]interface{}{
		"type":   "start",
		"symbol": 1,
		"vs":     p.Username,
	})

	p.Conn.WriteJSON(map[string]interface{}{
		"type":   "start",
		"symbol": 2,
		"vs":     waitingPlayer.Username,
	})

	// 🔄 Send initial board to BOTH
	for _, pl := range game.Players {
		if pl.Conn != nil {
			pl.Conn.WriteJSON(map[string]interface{}{
				"type":  "state",
				"board": game.Board,
				"turn":  game.Turn,
			})
		}
	}

	waitingPlayer = nil
	return game
}

func StartBotAfterTimeOut(p *Player) {
	time.Sleep(10 * time.Second)

	mutex.Lock()
	defer mutex.Unlock()

	if waitingPlayer != p {
		return
	}

	bot := &Player{
		Username: "BOT",
		Symbol:   2,
		Conn:     nil,
	}

	game := &Game{
		Turn:  1,
		Board: [ROWS][COLS]int{},
		Players: map[int]*Player{
			1: p,
			2: bot,
		},
	}

	p.Symbol = 1
	waitingPlayer = nil

	// 🔗 Attach game to player
	playerGames[p.Username] = game

	// 🔔 Notify client
	p.Conn.WriteJSON(map[string]interface{}{
		"type":   "start",
		"symbol": 1,
		"vs":     "BOT",
	})

	p.Conn.WriteJSON(map[string]interface{}{
		"type":  "state",
		"board": game.Board,
		"turn":  game.Turn,
	})
}
