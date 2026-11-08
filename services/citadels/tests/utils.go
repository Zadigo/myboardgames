package tests

import (
	"testing"

	"github.com/Zadigo/citadels/game"
	"github.com/Zadigo/citadels/internal/characters"
	"github.com/Zadigo/citadels/internal/districts"
	"github.com/Zadigo/citadels/internal/models"
)

// helper to create an empty target character (with no player, no gold)
func helperCreateEmptyTarget(t *testing.T) *characters.Architect {
	t.Helper()

	return &characters.Architect{}
}

// helper to create a character with a player
func helperCreateCharacter(name string, gold int) (models.PlayerInterface, models.CharacterInterface) {
	c, _ := characters.NewCharacter(name)

	p := models.NewPlayer()

	c.SetPlayer(p)
	p.IncreaseGold(gold)
	p.SetSelectedCharacter(c)

	return p, c
}

// helper to create a new game with players
func helperCreateGame(t *testing.T, players ...models.PlayerInterface) models.GameInterface {
	t.Helper()
	return &game.Game{
		Players: players,
	}
}

func helperCreatePlayers(n int) []models.PlayerInterface {
	players := make([]models.PlayerInterface, n)

	for i := range players {
		players[i] = models.NewPlayer()
	}

	return players
}

func helperAddDistrictsToPlayers(players []models.PlayerInterface) {
	for _, p := range players {
		p.AddDistrict(&districts.District{})
	}
}
