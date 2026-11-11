package tests

import (
	"testing"

	"github.com/Zadigo/citadels/internal/characters"
	"github.com/Zadigo/citadels/internal/game"
	"github.com/Zadigo/citadels/internal/models"
)

func TestSelectionPhase(t *testing.T) {
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
}
