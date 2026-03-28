package main

import "github.com/Zadigo/flipseven/cards"

type WebsocketMessage struct {
	Action  string `json:"action"`
	Message string `json:"message"`
	Deck    []cards.Card `json:"deck"`
}
