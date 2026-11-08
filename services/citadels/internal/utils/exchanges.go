package utils

import "github.com/Zadigo/citadels/internal/models"

// ExchangeDistricts swaps districts between two players, keeping order intact
func ExchangeDistricts(player models.PlayerInterface, withPlayer models.PlayerInterface) {
	playerDistricts := make([]models.DistrictInterface, len(player.GetDistricts()))
	targetDistricts := make([]models.DistrictInterface, len(withPlayer.GetDistricts()))

	for i, d := range withPlayer.GetDistricts() {
		targetDistricts[i] = d
	}

	for i, d := range player.GetDistricts() {
		playerDistricts[i] = d
	}

	player.ReplaceDistricts(targetDistricts)
	withPlayer.ReplaceDistricts(playerDistricts)
}
