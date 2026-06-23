package tests

import (
	"testing"

	"github.com/Zadigo/basegame/internal/game/cards"
	"github.com/stretchr/testify/assert"
)

func TestMakeCards(t *testing.T) {
	t.Run("Make Standard Cards", func(t *testing.T) {
		factory := cards.MakeCards(cards.STANDARD_CARD)
		items := factory.CreateCards()
		assert.Equal(t, 12, len(items), "Expected 12 standard cards")

		for _, card := range items {
			assert.Equal(t, cards.STANDARD_CARD, card.(*cards.BaseCard).CardType, "Expected card type to be STANDARD_CARD")
		}
	})

	t.Run("Make Standard Special Cards", func(t *testing.T) {
		factory := cards.MakeCards(cards.STANDARD_SPECIAL_CARD)
		items := factory.CreateCards()
		assert.Equal(t, 9, len(items), "Expected 9 standard special cards")
		for _, card := range items {
			assert.Equal(t, cards.STANDARD_SPECIAL_CARD, card.(*cards.BaseCard).CardType, "Expected card type to be STANDARD_SPECIAL_CARD")
			assert.True(t, card.(*cards.BaseCard).CardValue == 0, "Expected card value to be 0 for special cards")
		}
	})
}

func TestCardPile(t *testing.T) {
	pile := cards.NewCardPile()
	factory := cards.MakeCards(cards.STANDARD_CARD)
	items := factory.CreateCards()
	pile.AddCards(items...)

	card := pile.DrawCard("player-uuid")
	assert.NotNil(t, card, "Expected to draw a card from the pile")
	assert.Equal(t, 11, pile.RemainingCards(), "Expected 11 cards remaining in the pile after drawing one card")
	assert.NotEmpty(t, card.GetPlayer(), "Expected the card to have a player assigned")
}

func TestCardsRedis(t *testing.T) {
	redisClient := CreateRedisClient()
	defer redisClient.Close()

	client := cards.NewCardsRedis(t.Context(), redisClient)

	factory := cards.MakeCards(cards.STANDARD_CARD)
	items := factory.CreateCards()

	err := client.Save(items)
	assert.NoError(t, err, "Expected no error when saving cards to Redis")

	retrievedCards, err := client.Get()
	assert.NoError(t, err, "Expected no error when retrieving cards from Redis")
	assert.Equal(t, len(items), len(retrievedCards), "Expected the number of retrieved cards to match the number of saved cards")
}
