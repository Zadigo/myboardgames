package models

type GameInterface interface {
	GetPlayers() []Player
}

type BaseGame struct{}
