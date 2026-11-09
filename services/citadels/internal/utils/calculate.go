package utils

import (
	"github.com/Zadigo/citadels/internal/models"
)

type ScoringOptions struct {
	FinishedFirst bool
}

func calculateUniqueDistrictsPoints(player models.PlayerInterface) {
	for _, d := range player.GetDistricts() {
		d.Resolve()
	}
}

// CalculateScore calculates the player's score based on districts.
//   - Score points equal to the building cost of each of your districts
//   - If your city has at least one district of each type, score 3 points
//   - The player who first completed their city scores 4 points
//   - Any other player who completed their city scores 2 points
//   - Score any extra points from your unique districts
func CalculateScore(player models.PlayerInterface, options ScoringOptions) int {
	districtTypes := make(map[string]bool)

	score := 0

	for _, d := range player.GetDistricts() {
		score += d.GetCost()
		districtTypes[d.GetSymbol()] = true
	}

	var builtTypes int
	for _, v := range districtTypes {
		if v {
			builtTypes++
		}
	}

	if builtTypes == 5 {
		score += 3
	}

	if options.FinishedFirst {
		score += 4
	}

	calculateUniqueDistrictsPoints(player)
	player.AddPoints(score)
	return score
}
