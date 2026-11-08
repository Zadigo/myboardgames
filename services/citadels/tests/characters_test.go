package tests

import (
	"testing"

	"github.com/Zadigo/citadels/internal/characters"
	"github.com/Zadigo/citadels/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestBaseCharacter(t *testing.T) {
	char, err := characters.NewCharacter("assassin")

	assert.NotNil(t, char)
	assert.NoError(t, err)
	assert.Equal(t, "assassin", char.GetName())
	assert.Equal(t, 1, char.GetLevel())

	assert.False(t, char.IsKilled())

	t.Run("resolve base character", func(t *testing.T) {
		simpleTarget := &characters.Character{Name: "target"}
		assert.NoError(t, char.Resolve(models.CharacterOptions{Target: simpleTarget}))
	})
}

func TestAssassin(t *testing.T) {
	t.Run("with valid target", func(t *testing.T) {
		char := characters.Assassin{
			BaseCharacter: models.BaseCharacter{},
		}

		target := characters.Architect{
			BaseCharacter: models.BaseCharacter{},
		}

		err := char.Resolve(models.CharacterOptions{Target: &target})
		assert.NoError(t, err)
		assert.True(t, target.IsKilled())
	})

	t.Run("with invalid target", func(t *testing.T) {
		char := characters.Assassin{}
		err := char.Resolve(models.CharacterOptions{Target: nil})
		assert.Error(t, err)
	})
}
