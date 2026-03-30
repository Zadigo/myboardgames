package cards

import (
	"github.com/Zadigo/flipseven/internal"
)

// Checks the cards for the given player. A player cannot have
// two cards with the same number (except for special and bonus cards)
func CheckPlayerCards(player *internal.Player) {
	seen := make(map[int]bool)

	for _, card := range player.Cards {
		if card.IsBonus || card.IsMultiplier {
			continue
		}

		if seen[card.Value] {
			player.IsFreezed = true
			return
		}
		seen[card.Value] = true
	}

	if len(player.Cards) == 7 {
		player.HasSevenCards = true
		player.IsFreezed = true
	}
}

// Once a card is flipped from the deck, attributes it to a given player
func AttributeCardToPlayer(card internal.Card, player *internal.ConnectedPlayer) {
	player.Details.Cards = append(player.Details.Cards, card)
	player.Details.NumberOfCards += 1
	CheckPlayerCards(&player.Details)
}

// Pops x number of cards from the deck
func FlipCard(deck []internal.Card, k int) internal.Card {
	return deck[k]
}
