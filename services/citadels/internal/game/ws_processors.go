package game

import (
	"github.com/Zadigo/citadels/internal/characters"
	"github.com/Zadigo/citadels/internal/models"
	"github.com/gorilla/websocket"
)

func SelectionPhase(conn *websocket.Conn, game *models.BaseGame, message models.WebsocketMessage) {
	switch message.Action {
	case models.ApplyCharacterSelection:
		characters, _ := characters.NewCharacters(message.CharacterSelection.Names...)
		CreateIndicators(game.Table, characters)
	}
}

func TurnPhase(conn *websocket.Conn, game *models.BaseGame, message models.WebsocketMessage) {
}
