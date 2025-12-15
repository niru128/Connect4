package main

import (
	"encoding/json"
	"net/http"
)

func HandleLeaderboard(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")
	leaderboardMutex.Lock()
	defer leaderboardMutex.Unlock()

	var result []LeaderboardEntry

	for user, wins := range leaderboard {
		result = append(result, LeaderboardEntry{
			Username: user,
			Wins:     wins,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
