package logic

func SumValues(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Calculate the player's score by summing the values
// of the cards in their hand
func (player *Player) TotalPoints() int {
	var normalPoints int = 0
	var multiplierValue int = 0

	for _, card := range player.Cards {
		if card.IsNumber || card.IsBonus {
			// Calculate normal points
			normalPoints += card.Value
		}

		if card.IsMultiplier {
			multiplierValue += card.Value
		}
	}

	// Calculate final score
	if multiplierValue > 0 {
		player.Score = normalPoints * multiplierValue
	} else {
		player.Score = normalPoints
	}

	return player.Score
}

// Reset the player's state for a new round
func (player *PlayerLayer) ResetPlayerState() {
	player.Details.IsFreezed = false
	player.Details.HasSevenCards = false
	player.Details.HasSecondChance = false
	player.Details.Cards = []Card{}
}

// Number of cards that the player has flipped
func (player *PlayerLayer) NumberOfCards() int {
	return len(player.Details.Cards)
}
