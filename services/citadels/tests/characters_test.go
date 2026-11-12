package tests

import (
	"testing"

	"github.com/Zadigo/citadels/internal/characters"
	"github.com/Zadigo/citadels/internal/districts"
	"github.com/Zadigo/citadels/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestNewCharacterFunctions(t *testing.T) {
	t.Run("create a single character", func(t *testing.T) {
		char, err := characters.NewCharacter(characters.KingCharacter)
		assert.NoError(t, err)
		assert.NotNil(t, char)
	})

	t.Run("create new characters", func(t *testing.T) {
		chars, err := characters.NewCharacters(characters.AssassinCharacter, characters.ThiefCharacter)

		assert.NoError(t, err)
		assert.Len(t, chars, 2)
		assert.IsType(t, &characters.Assassin{}, chars[0])
		assert.IsType(t, &characters.Thief{}, chars[1])
	})
}

func TestBaseCharacter(t *testing.T) {
	char := &characters.Character{}

	assert.NotNil(t, char)
	assert.False(t, char.IsKilled())

	t.Run("resolve base character", func(t *testing.T) {
		simpleTarget := &characters.Character{Name: "target"}
		assert.NoError(t, char.Resolve(models.CharacterOptions{Target: simpleTarget}))
	})

	t.Run("should reset correctly", func(t *testing.T) {
		t.Parallel()

		other := &characters.Character{
			CanBuildDistricts:        false,
			CanUseCharacterAbilities: false,
			Bewitched:                &characters.Witch{},
			Player:                   &models.Player{},
			Killed:                   true,
		}

		other.Reset()

		assert.True(t, other.CanBuildDistricts)
		assert.True(t, other.CanUseCharacterAbilities)
		assert.False(t, other.Killed)
		assert.Nil(t, other.Bewitched)
		assert.Nil(t, other.Player)
	})

	t.Run("should not betwitch if witch", func(t *testing.T) {
		t.Parallel()

		w := &characters.Character{Name: "witch"}

		err := w.BeWitch(&characters.Witch{})
		assert.Error(t, err)
	})

	t.Run("should bewitch if witch", func(t *testing.T) {
		t.Parallel()

		a := &characters.Character{Name: "witch"}
		other := &characters.Character{}
		other.BeWitch(a)

		assert.NotNil(t, other.Bewitched)
		assert.False(t, other.CanBuildDistricts)
		assert.False(t, other.CanUseCharacterAbilities)

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

	t.Run("should fail", func(t *testing.T) {
		t.Run("attacking self", func(t *testing.T) {
			t.Parallel()

			players := helperCreatePlayers(2)

			a := &characters.Thief{Player: players[0]}
			b := &characters.Thief{Player: players[0]}

			err := a.Resolve(models.CharacterOptions{Target: b})

			assert.Error(t, err)
			assert.EqualError(t, err, characters.CannotAttackSelfError.Error())
		})

		// t.Run("target has no selected character", func(t *testing.T) {
		// 	t.Parallel()

		// 	p1, attacker := helperCreateCharacter("thief", 0)
		// 	p2, target := helperCreateCharacter("architect", 0)

		// 	target.GetPlayer().SetSelectedCharacter(nil)

		// 	game := helperCreateGame(t, p1, p2)
		// 	err := attacker.Resolve(models.CharacterOptions{Target: target, CurrentGame: game})
		// 	assert.Error(t, err)
		// 	assert.EqualError(t, err, characters.TargetHasNoCharacterError.Error())
		// })

		t.Run("target is level 1", func(t *testing.T) {
			t.Parallel()

			p1, attacker := helperCreateCharacter("thief", 0)
			p2, target := helperCreateCharacter("assassin", 10)
			game := helperCreateGame(t, p1, p2)

			target.Kill()

			err := attacker.Resolve(models.CharacterOptions{Target: target, CurrentGame: game})

			assert.Error(t, err)
		})

		t.Run("target is killed", func(t *testing.T) {
			t.Parallel()

			p1, attacker := helperCreateCharacter("thief", 0)
			p2, target := helperCreateCharacter("architect", 10)
			game := helperCreateGame(t, p1, p2)

			target.Kill()

			err := attacker.Resolve(models.CharacterOptions{Target: target, CurrentGame: game})

			assert.Error(t, err)
		})

		t.Run("target is bewitched", func(t *testing.T) {
			t.Parallel()

			players := helperCreatePlayers(3)

			c := &characters.Witch{Player: players[2]}
			a := &characters.Thief{Player: players[0]}
			b := &characters.Architect{Player: players[1], Bewitched: c}

			err := a.Resolve(models.CharacterOptions{Target: b})
			assert.Error(t, err)
			assert.EqualError(t, err, characters.CannotStealBewitchedError.Error())
		})
	})

	t.Run("target was robbed successfully", func(t *testing.T) {
		players := helperCreatePlayers(2)

		a := &characters.Thief{Player: players[0]}
		b := &characters.Architect{Player: players[1]}

		b.GetPlayer().IncreaseGold(10)

		err := a.Resolve(models.CharacterOptions{Target: b})
		assert.NoError(t, err)
		assert.Equal(t, 0, b.GetPlayer().GoldCount())
		assert.Equal(t, 14, a.GetPlayer().GoldCount())
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

func TestMagistrate(t *testing.T) {
	t.Run("cannot place warrant on self", func(t *testing.T) {})

	t.Run("can place warrant on other characters", func(t *testing.T) {
		players := helperCreatePlayers(3)

		m := &characters.Magistrate{Player: players[0]}
		a := &characters.Architect{Player: players[1]}
		b := &characters.Scholar{Player: players[2]}
		c := &characters.Blackmailer{Player: players[2]}

		a.PlaceSignedWarrant(m)
		b.PlaceUnsignedWarrant(m)
		c.PlaceWarrantMarkerBack(m)

		assert.True(t, a.SignedWarrant)
		assert.True(t, b.UnsignedWarrant)
		assert.True(t, c.WarrantMarkerBack)
	})
}

func TestMerchant(t *testing.T) {
	t.Run("should add points for trade district", func(t *testing.T) {
		t.Parallel()

		players := helperCreatePlayers(1)
		p1 := players[0]

		d1 := &districts.District{DistrictType: districts.Trade}
		p1.AddDistrict(d1)

		merchant := &characters.Merchant{Player: p1}
		merchant.Resolve(models.CharacterOptions{})

		assert.Equal(t, 1, p1.GetPoints())
		assert.Equal(t, 3, p1.GoldCount())
	})

	t.Run("no trade district", func(t *testing.T) {
		t.Parallel()

		players := helperCreatePlayers(1)
		p1 := players[0]

		d1 := &districts.District{DistrictType: districts.Nobel}
		p1.AddDistrict(d1)

		merchant := characters.Merchant{Player: p1}
		merchant.Resolve(models.CharacterOptions{})

		assert.Equal(t, 0, p1.GetPoints())
		assert.Equal(t, 3, p1.GoldCount())
	})
}

func TestNavigator(t *testing.T) {
	t.Run("should gain four golds", func(t *testing.T) {
		t.Parallel()

		players := helperCreatePlayers(1)
		p1 := players[0]

		navigator := characters.Navigator{Player: p1}
		err := navigator.Resolve(models.CharacterOptions{
			Navigator: models.NavigatorOptions{
				GainFourGolds: true,
			},
		})

		assert.NoError(t, err)
		assert.Equal(t, p1.GoldCount(), 6)
	})

	t.Run("should gain four cards", func(t *testing.T) {
		t.Skip("Not implemented yet")
		t.Parallel()

		players := helperCreatePlayers(1)
		p1 := players[0]

		d1 := &districts.District{DistrictType: districts.Nobel}
		p1.AddDistrict(d1)

		navigator := characters.Navigator{Player: p1}
		err := navigator.Resolve(models.CharacterOptions{
			Navigator: models.NavigatorOptions{
				GainFourCards: true,
			},
		})

		assert.NoError(t, err)
		assert.Equal(t, p1.GetPoints(), 1)
	})

	t.Run("should raise errors", func(t *testing.T) {
		t.Parallel()

		t.Run("both options true", func(t *testing.T) {
			navigator := characters.Navigator{}
			err := navigator.Resolve(models.CharacterOptions{
				Navigator: models.NavigatorOptions{
					GainFourGolds: true,
					GainFourCards: true,
				},
			})
			assert.Error(t, err)
		})

		t.Run("both options false", func(t *testing.T) {
			navigator := characters.Navigator{}
			err := navigator.Resolve(models.CharacterOptions{
				Navigator: models.NavigatorOptions{},
			})
			assert.Error(t, err)
		})
	})
}
