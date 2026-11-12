package utils

import "github.com/Zadigo/citadels/internal/models"

// ExchangeDistricts swaps districts between two players, keeping order intact
func ExchangeDistricts(player models.PlayerInterface, withPlayer models.PlayerInterface) {
	playerDistricts := make([]models.DistrictInterface, len(player.GetBuiltDistricts()))
	targetDistricts := make([]models.DistrictInterface, len(withPlayer.GetBuiltDistricts()))

	for i, d := range withPlayer.GetBuiltDistricts() {
		targetDistricts[i] = d
	}

	for i, d := range player.GetBuiltDistricts() {
		playerDistricts[i] = d
	}

	player.ReplaceDistricts(targetDistricts)
	withPlayer.ReplaceDistricts(playerDistricts)
}
