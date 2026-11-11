package game

import (
	"github.com/Zadigo/citadels/internal/characters"
	"github.com/Zadigo/citadels/internal/districts"
	"github.com/Zadigo/citadels/internal/models"
	"github.com/gorilla/websocket"
)

func SelectionPhase(conn *websocket.Conn, gameInstance models.GameInterface, message models.WebsocketMessage) {
	switch message.Action {
	case models.ApplyCharacterSelection:
		// 1. Character selection phase
		characters, _ := characters.NewCharacters(message.CharacterSelection.Names...)
		_ = CreateIndicators(gameInstance.GetTable(), characters)

		// 2. Create unique districts + base districts
		_, _ = districts.ReadYamlFile()
	case models.AcceptingNewPlayer:
		// maybe add player to game
	}
}

func TurnPhase(conn *websocket.Conn, gameInstance *models.BaseGame, message models.WebsocketMessage) {
}
