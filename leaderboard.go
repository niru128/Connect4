package main

import "sync"

var leaderboard = make(map[string]int)
var leaderboardMutex sync.Mutex

func RecordWin(username string) {
	leaderboardMutex.Lock()
	defer leaderboardMutex.Unlock()
	leaderboard[username]++
}

type LeaderboardEntry struct {
	Username string `json:"username"`
	Wins     int    `json:"wins"`
}
