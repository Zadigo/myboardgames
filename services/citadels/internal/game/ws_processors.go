package game

import (
	"fmt"

	"github.com/Zadigo/citadels/internal/characters"
	"github.com/Zadigo/citadels/internal/districts"
	"github.com/Zadigo/citadels/internal/models"
	"github.com/gorilla/websocket"
)

func CreationPhase(conn *websocket.Conn, message models.WebsocketMessage) {
	switch message.Action {
	case models.NewGame:
		gameInstance := NewGame()
		conn.WriteJSON(models.WebsocketMessage{
			Action:  models.NewGame,
			Message: gameInstance.GetID(),
		})
	}
}

func WaitingPhase(conn *websocket.Conn, gameInstance models.GameInterface, message models.WebsocketMessage) {
	switch message.Action {
	case models.AcceptingNewPlayer:
		player := models.NewPlayer(conn)
		remaining := gameInstance.GetTable().AddPlayer(player)
		conn.WriteJSON(models.WebsocketMessage{Message: fmt.Sprintf("Remaining %d spots", remaining)})
	}
}

func SelectionPhase(conn *websocket.Conn, gameInstance models.GameInterface, message models.WebsocketMessage) {
	switch message.Action {
	case models.ApplyCharacterSelection:
		// 1. Character selection phase
		characters, _ := characters.NewCharacters(message.CharacterSelection.Names...)
		_ = CreateIndicators(gameInstance.GetTable(), characters)

		// 2. Create unique districts + base districts
		_, _ = districts.ReadYamlFile()
	}
}

func TurnPhase(conn *websocket.Conn, gameInstance *models.BaseGame, message models.WebsocketMessage) {
}
