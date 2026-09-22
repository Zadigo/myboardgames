package cards

import (
	"github.com/Zadigo/basegame/internal/models"
	"github.com/google/uuid"
)

type StandardSpecialCardsFactory struct {
}

func (j *StandardSpecialCardsFactory) CreateCards() []models.CardInterface {
	names := []string{OPERATION_FREEZE, OPERATION_FLIP3, OPERATION_SECOND_CHANCE}
	cards := []models.CardInterface{}

	for _, name := range names {
		for range 3 {
			cards = append(cards, &BaseCard{
				Uuid:         uuid.NewString(),
				CardType:     STANDARD_SPECIAL_CARD,
				CardValue:    0,
				CardOperator: name,
			})
		}
	}

	return cards
}
