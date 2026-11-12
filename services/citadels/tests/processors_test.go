package tests

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/Zadigo/citadels/internal/characters"
	"github.com/Zadigo/citadels/internal/game"
	"github.com/Zadigo/citadels/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestWaitingPhase(t *testing.T) {
	t.Run("should add player", func(t *testing.T) {
		t.Parallel()

		conn, _, _, _ := helperCreateConnection(t, "/new")
		instance := game.NewGame()
		game.WaitingPhase(conn, instance, models.WebsocketMessage{
			Action: models.AcceptingNewPlayer,
		})

		assert.Len(t, instance.GetTable().Players, 1)

		response := map[string]any{}
		err := conn.ReadJSON(&response)
		assert.NoError(t, err)

		log.Print(json.MarshalIndent(response, "", "  "))
	})
}

func TestSelectionPhase(t *testing.T) {
	t.Run("should select character", func(t *testing.T) {
		t.Parallel()

		conn, _, _, _ := helperCreateConnection(t, "/new")

		// 1. Character selection phase
		instance := game.NewGame()
		game.SelectionPhase(conn, instance, models.WebsocketMessage{
			Action: models.ApplyCharacterSelection,
			CharacterSelection: models.CharacterSelection{
				Names: []string{characters.ArchitectCharacter},
			},
		})
	})
}
