package cards

import "github.com/google/uuid"

type StandardSpecialCardsFactory struct {
}

func (j *StandardSpecialCardsFactory) createCards() []CardInterface {
	names := []string{OPERATION_FREEZE, OPERATION_FLIP3, OPERATION_SECOND_CHANCE}
	cards := []CardInterface{}

	for _, name := range names {
		cards = append(cards, &BaseCard{
			Uuid:         uuid.NewString(),
			CardType:     STANDARD_SPECIAL_CARD,
			CardValue:    0,
			CardOperator: name,
		})
	}

	return cards
}
