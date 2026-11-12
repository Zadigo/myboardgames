package tests

import (
	"testing"

	"github.com/Zadigo/citadels/internal/characters"
	"github.com/Zadigo/citadels/internal/game"
	"github.com/stretchr/testify/assert"
)

func TestTableManualSetup(t *testing.T) {
	t.Run("test table setup", func(t *testing.T) {
		game := game.NewGame()
		table := game.GetTable()

		assert.NotNil(t, table)

		players := helperCreatePlayers(4)

		for _, p := range players {
			table.AddPlayer(p)
		}

		testCharacters := []string{
			characters.AbbotCharacter,
			characters.ThiefCharacter,
			characters.AssassinCharacter,
			characters.ArtistCharacter,
			characters.KingCharacter,
			characters.TraderCharacter,
			characters.AlchemistCharacter,
			characters.BlackmailerCharacter,
		}

		testCharactersInterfaces, _ := characters.NewCharacters(testCharacters...)

		// Indicators
		t.Run("should create indicators", func(t *testing.T) {
			table.CreateCharacterIndicators(testCharactersInterfaces)
			// table.SelectCharacterCards(testCharacters...)

			indicators := table.GetCharacterIndicators()
			assert.Equal(t, 1, indicators[0].Level)
			assert.Greater(t, indicators[len(indicators)-1].Level, indicators[0].Level)
		})

		// Card distribution
		t.Run("should distribute districts", func(t *testing.T) {
			table.DistributeDistricts(4)
			for _, p := range table.Players {
				assert.Equal(t, 4, p.GetCurrentHand())
			}
		})

		// District pile

		// Give crown
		t.Run("should give crown", func(t *testing.T) {
			table.GiveCrown("random")
			p := table.GetPlayerWithCrown()
			assert.NotNil(t, p)
		})

		t.Run("player with crown should have characters to select from", func(t *testing.T) {
			table.StartRound()
			_ = table.GetPlayerWithCrown()
			// hand := p.GetCurrentCharacterHand()
			hand := table.GetCurrentCharacterHand()
			assert.NotNil(t, hand)
			assert.Len(t, hand, 8)
		})

		// State
		t.Run("state should be true", func(t *testing.T) {
			assert.True(t, table.GetState())
		})
	})
}

func BenchmarkTableSetup(b *testing.B) {
	b.ResetTimer()
}

// func TestTableCharacterResolution(t *testing.T) {
// 	t.Run("should resolve characters sequentially", func(t *testing.T) {
// 		game := game.NewGame()
// 		table := game.GetTable()

// 		table.Setup()
// 		assert.NotNil(t, table)

// 		table.StartRound()

// 		t.Run("every player should select a card", func(t *testing.T) {
// 			for _, p := range table.Players {
// 				c := table.GetCurrentPlayer()
// 				assert.NotNil(t, c)
// 				assert.Equal(t, p, c)

// 				table.CurrentPlayerSelectsCard("card-id")
// 				time.Sleep(1 * time.Second)
// 			}
// 		})

// 		t.Run("should resolve characters", func(t *testing.T) {
// 			for range 8 {
// 				r := table.ResolveCharacters()
// 				if r != nil {
// 					assert.Contains(t, c, r.Name)
// 				}
// 				r.Resolve(models.CharacterOptions{})
// 				time.Sleep(1 * time.Second)
// 			}
// 		})
// 	})
// }
