package cards

import "github.com/Zadigo/flipseven/internal"

func getNumberCards() []internal.Card {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	cards := []internal.Card{}

	for _, n := range numbers {
		if n <= 1 {
			// 0 and 1 appear once
			cards = append(cards, internal.Card{Value: n, IsNumber: true})
		} else {
			for j := range n {
				cards = append(cards, internal.Card{
					Value:    j,
					IsNumber: true,
				})
			}
		}
	}

	return cards
}

func getBonusCards() []internal.Card {
	cards := []internal.Card{}
	numbers := []int{2, 4, 6, 8, 10}

	cards = append(cards, internal.Card{
		Value:        2,
		IsMultiplier: true,
		IsSpecial:    true,
	})

	for i := range numbers {
		cards = append(cards, internal.Card{
			Value:        i,
			IsBonus:      true,
			IsMultiplier: false,
		})
	}

	return cards
}

func getSpecialCards() []internal.Card {
	cards := []internal.Card{}
	names := []string{"Freeze", "Flip 3", "Second Chance"}

	for i := range names {
		for range 3 {
			cards = append(cards, internal.Card{
				Value:    0,
				Category: names[i],
			})
		}
	}

	return cards
}

// Create a new deck of X cards that
// will serve as the main deck for
// flipping cards for each player
func GetDeck() []internal.Card {
	cards := []internal.Card{}

	numberCards := getNumberCards()
	cards = append(cards, numberCards...)

	bonuCards := getBonusCards()
	cards = append(cards, bonuCards...)

	specialCards := getSpecialCards()
	cards = append(cards, specialCards...)

	return cards
}

func GetShuffledDeck() []internal.Card {
	deck := GetDeck()

	copiedDeck := make([]internal.Card, len(deck))
	copy(copiedDeck, deck)

	return copiedDeck
}
