package game

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/Zadigo/citadels/internal/models"
)

// func CreateIndicators(names []models.CharacterInterface) []models.BaseCharacterIndicator {
// 	indicators := make([]models.BaseCharacterIndicator, len(names))
// 	for i, n := range names {
// 		indicators[i] = models.BaseCharacterIndicator{
// 			Name:  n.GetName(),
// 			Level: n.GetLevel(),
// 		}
// 	}

// 	slices.SortFunc(indicators, func(a, b models.BaseCharacterIndicator) int {
// 		return cmp.Compare(a.Level, b.Level)
// 	})

// 	return indicators
// }

func CreateIndicators(table *models.BaseTable, names []models.CharacterInterface) error {
	if len(names) > 8 {
		return fmt.Errorf("Can only have 8 characters at most")
	}

	for i, n := range names {
		table.ResolutionIndicator[i] = models.BaseCharacterIndicator{
			Name:  n.GetName(),
			Level: n.GetLevel(),
		}
	}

	slices.SortFunc(table.ResolutionIndicator, func(a, b models.BaseCharacterIndicator) int {
		return cmp.Compare(a.Level, b.Level)
	})

	return nil
}
