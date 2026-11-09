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
		t.Parallel()

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

	t.Run("should be ReplacedTarget when set", func(t *testing.T) {
		t.Parallel()

		char := characters.Assassin{}
		target := characters.Architect{}

		err := char.Resolve(models.CharacterOptions{ReplacedTarget: &target})

		assert.NoError(t, err)
		assert.True(t, target.IsKilled())
	})

	t.Run("should raise errors", func(t *testing.T) {
		t.Run("with invalid targets", func(t *testing.T) {
			t.Parallel()

			char := characters.Assassin{}
			err := char.Resolve(models.CharacterOptions{Target: nil, ReplacedTarget: nil})
			assert.Error(t, err)
			assert.EqualError(t, err, characters.MissingTargetError.Error())
		})

		t.Run("ambiguous target", func(t *testing.T) {
			t.Parallel()

			w := characters.Assassin{}
			err := w.Resolve(models.CharacterOptions{Target: &characters.Assassin{}, ReplacedTarget: &characters.Architect{}})
			assert.Error(t, err)
			assert.EqualError(t, err, characters.AmbiguousTargetError.Error())
		})
	})
}

func TestThief(t *testing.T) {
	t.Run("target has no player", func(t *testing.T) {
		target := helperCreateEmptyTarget(t)
		_, attacker := helperCreateCharacter("thief", 0)

		err := attacker.Resolve(models.CharacterOptions{Target: target})
		assert.Error(t, err)
	})

	t.Run("attacking self", func(t *testing.T) {
		p1, attacker := helperCreateCharacter("thief", 0)
		game := helperCreateGame(t, p1)

		err := attacker.Resolve(models.CharacterOptions{Target: attacker, CurrentGame: game})

		assert.Error(t, err)
		assert.EqualError(t, err, characters.CannotAttackSelfError.Error())
	})

	t.Run("target has no selected character", func(t *testing.T) {
		p1, attacker := helperCreateCharacter("thief", 0)
		p2, target := helperCreateCharacter("architect", 0)

		target.GetPlayer().SetSelectedCharacter(nil)

		game := helperCreateGame(t, p1, p2)
		err := attacker.Resolve(models.CharacterOptions{Target: target, CurrentGame: game})
		assert.Error(t, err)
		assert.EqualError(t, err, characters.TargetHasNoCharacterError.Error())
	})

	t.Run("target is level 1", func(t *testing.T) {
		p1, attacker := helperCreateCharacter("thief", 0)
		p2, target := helperCreateCharacter("assassin", 10)
		game := helperCreateGame(t, p1, p2)

		target.Kill()

		err := attacker.Resolve(models.CharacterOptions{Target: target, CurrentGame: game})

		assert.Error(t, err)
	})

	t.Run("target is killed", func(t *testing.T) {
		p1, attacker := helperCreateCharacter("thief", 0)
		p2, target := helperCreateCharacter("architect", 10)
		game := helperCreateGame(t, p1, p2)

		target.Kill()

		err := attacker.Resolve(models.CharacterOptions{Target: target, CurrentGame: game})

		assert.Error(t, err)
	})

	t.Run("target is bewitched", func(t *testing.T) {})

	t.Run("target was robbed successfully", func(t *testing.T) {
		p1, attacker := helperCreateCharacter("thief", 0)
		p2, target := helperCreateCharacter("architect", 10)
		game := helperCreateGame(t, p1, p2)

		err := attacker.Resolve(models.CharacterOptions{Target: target, CurrentGame: game})
		assert.NoError(t, err)
		assert.Equal(t, 0, target.GetPlayer().GoldCount())
		assert.Equal(t, 10, attacker.GetPlayer().GoldCount())
	})
}

func TestWitch(t *testing.T) {
	players := helperCreatePlayers(3)

	t.Run("should bewitch character", func(t *testing.T) {
		a := characters.Assassin{
			Player: players[0],
		}

		w := characters.Witch{
			Player: players[1],
		}

		z := characters.Architect{
			Player: players[2],
		}

		err := w.Resolve(models.CharacterOptions{
			Target:         &a,
			ReplacedTarget: &z,
		})

		assert.NoError(t, err)
		assert.True(t, a.IsBewitched())
		assert.False(t, a.CanBuildDistricts)
		assert.False(t, a.CanUseCharacterAbilities)
		assert.True(t, z.IsKilled())
	})
}
