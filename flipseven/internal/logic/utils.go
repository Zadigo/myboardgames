package logic

import (
	"github.com/Zadigo/flipseven/internal"
)

func SumValues(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Calculate the player's score by summing the values
// of the cards in their hand
func CalculatePlayerScores(player *internal.ConnectedPlayer) {
	var normalPoints int = 0
	var multipluerValue int = 0

	for _, card := range player.Details.Cards {
		if card.IsNumber || card.IsBonus {
			// Calculate normal points
			normalPoints += card.Value
		}

		if card.IsMultiplier {
			multipluerValue += card.Value
		}
	}

	// Calculate final score
	if multipluerValue > 0 {
		player.Details.Score = normalPoints * multipluerValue
	} else {
		player.Details.Score = normalPoints
	}
}

// Reset the player's state for a new round
func RestPlayerState(player *internal.ConnectedPlayer) {
	player.Details.NumberOfCards = 0
	player.Details.IsFreezed = false
	player.Details.HasSevenCards = false
	player.Details.HasSecondChance = false
	player.Details.Cards = []internal.Card{}
}
