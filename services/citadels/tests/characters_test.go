package tests

import (
	"testing"

	"github.com/Zadigo/citadels/internal/characters"
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
		simpleTarget := &characters.BaseCharacter{Name: "target"}
		assert.NoError(t, char.Resolve(characters.CharacterOptions{Target: simpleTarget}))
	})
}

func TestAssassin(t *testing.T) {
	t.Run("with valid target", func(t *testing.T) {
		char := characters.Assassin{
			BaseCharacter: characters.BaseCharacter{},
		}

		target := characters.Architect{
			BaseCharacter: characters.BaseCharacter{},
		}

		err := char.Resolve(characters.CharacterOptions{Target: &target})
		assert.NoError(t, err)
		assert.True(t, target.IsKilled())
	})

	t.Run("with invalid target", func(t *testing.T) {
		char := characters.Assassin{}
		err := char.Resolve(characters.CharacterOptions{Target: nil})
		assert.Error(t, err)
	})
}
