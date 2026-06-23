package cards

import "github.com/google/uuid"

type StandardCardFactory struct {
}

func (s *StandardCardFactory) CreateCards() []CardInterface {
	cards := []CardInterface{}

	for i := range 12 {
		cards = append(cards, &BaseCard{
			Uuid:         uuid.NewString(),
			CardType:     STANDARD_CARD,
			CardValue:    i,
			CardOperator: OPERATION_ADD,
		})
	}

	return cards
}
