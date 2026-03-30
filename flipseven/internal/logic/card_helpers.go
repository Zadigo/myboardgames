package logic

// // Checks the cards for the given player. A player cannot have
// // two cards with the same number (except for special and bonus cards)
// func (player *ConnectedPlayer) CheckPlayerCards() {
// 	seen := make(map[int]bool)

// 	for _, card := range player.Details.Cards {
// 		if card.IsBonus || card.IsMultiplier {
// 			continue
// 		}

// 		if seen[card.Value] {
// 			player.Details.IsFreezed = true
// 			return
// 		}
// 		seen[card.Value] = true
// 	}

// 	if len(player.Details.Cards) == 7 {
// 		player.Details.HasSevenCards = true
// 		player.Details.IsFreezed = true
// 	}
// }

// // Once a card is flipped from the deck, attributes it to a given player
// func (player *ConnectedPlayer) AttributeCardToPlayer(card Card) {
// 	player.Details.Cards = append(player.Details.Cards, card)
// 	player.Details.NumberOfCards += 1
// 	player.CheckPlayerCards()
// }
