package tests

import (
	"testing"

	"github.com/Zadigo/citadels/internal/characters"
	"github.com/Zadigo/citadels/internal/game"
	"github.com/Zadigo/citadels/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	t.Run("should instantiate game", func(t *testing.T) {
		instance := game.NewGame()

		assert.IsType(t, instance, &game.Game{})
		assert.NotEmpty(t, instance.GetID())
		assert.Len(t, instance.GetPlayers(), 8)
		assert.Len(t, instance.GetCharacters(), 8)
		assert.Nil(t, instance.GetCrownOwner())
		assert.Equal(t, instance.GetStage(), models.WaitingForPlayers)
	})

	t.Run("should add player when below 8", func(t *testing.T) {
		t.Parallel()

		players := helperCreatePlayers(8)
		instance := game.NewGame()

		remainders := make([]int, 0)
		for _, player := range players {
			rem := instance.GetTable().AddPlayer(player)
			remainders = append(remainders, rem)
			assert.IsDecreasing(t, remainders, rem)
		}
	})
}

func TestCreateIndicators(t *testing.T) {
	t.Run("should create indicators", func(t *testing.T) {
		characters := []models.CharacterInterface{
			&characters.Character{Name: "A", Level: 8},
			&characters.Character{Name: "A", Level: 7},
			&characters.Character{Name: "A", Level: 6},
			&characters.Character{Name: "A", Level: 5},
			&characters.Character{Name: "A", Level: 4},
			&characters.Character{Name: "A", Level: 3},
			&characters.Character{Name: "B", Level: 2},
			&characters.Character{Name: "C", Level: 1},
		}

		table := &models.BaseTable{ResolutionIndicator: make([]models.BaseCharacterIndicator, 8)}
		err := game.CreateIndicators(table, characters)
		assert.NoError(t, err)
	})
}
