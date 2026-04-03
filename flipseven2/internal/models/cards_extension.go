package models

func GetExtensionSpecialCards() []Card {
	cards := []Card{}
	names := []string{}

	for i := range names {
		for range 3 {
			cards = append(cards, Card{
				Value:    0,
				Category: names[i],
			})
		}
	}

	return cards
}
