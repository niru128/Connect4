package main

import "github.com/gorilla/websocket"

type Player struct {
	ID       string
	Username string
	Conn     *websocket.Conn
	Symbol   int
}

type WSMessage struct {
	Type   string `json:"type"`
	Column int    `json:"column,omitempty"`
}

type Bot struct {
	Symbol int
}
