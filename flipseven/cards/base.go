package cards


func getNumberCards() []Card {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	cards := []Card{}

	for i := range numbers {
		if i == 0 {
			cards = append(cards, Card{
				value:    0,
				isNumber: true,
			})
		}

		if i == 1 {
			cards = append(cards, Card{
				value:    0,
				isNumber: true,
			})
		}

		if i > 1 {
			for j := range i {
				cards = append(cards, Card{
					value:    j,
					isNumber: true,
				})
			}
		}
	}

	return cards
}

func getBonusCards() []Card {
	cards := []Card{}
	numbers := []int{2, 4, 6, 8, 10}

	cards = append(cards, Card{
		value:        2,
		isMultiplier: true,
		isSpecial:    true,
	})

	for i := range numbers {
		cards = append(cards, Card{
			value:        i,
			isBonus:      true,
			isMultiplier: false,
		})
	}

	return cards
}

func getSpecialCards() []Card {
	cards := []Card{}
	names := []string{"Freeze", "Flip 3", "Second Chance"}

	for i := range names {
		for range 3 {
			cards = append(cards, Card{
				value:    0,
				category: names[i],
			})
		}
	}

	return cards
}

// Create a new deck of X cards that
// will serve as the main deck for
// flipping cards for each player
func GetDeck() []Card {
	cards := []Card{}

	numberCards := getNumberCards()
	cards = append(cards, numberCards...)

	bonuCards := getBonusCards()
	cards = append(cards, bonuCards...)

	specialCards := getSpecialCards()
	cards = append(cards, specialCards...)

	return cards
}
