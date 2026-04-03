package models

type Card struct {
	// The card's value, for number cards this is the number, for multiplier
	// cards this is the multiplier, for bonus cards this is the bonus points,
	// for special cards this is the special effect
	Value int `json:"value"`
	// The player who owns the card, nil for unowned
	Owner *Player `json:"owner"`
	// nil, Freeze, Flip 3 or Second Chance
	Category string `json:"category"`
	// Is x2 instead of +2
	IsMultiplier bool `json:"isMultiplier"`
	// Base card
	IsNumber bool `json:"isNumber"`
	// One of x2, +2, +4, +6, +8 or +10
	IsBonus bool `json:"isBonus"`
	// One of: Freeze, Flip 3 or Second Chance
	IsSpecial bool `json:"isSpecial"`
}

func (c *Card) SetOwner(owner *Player) {
	c.Owner = owner
}

// GetNumberCards returns a slice of number cards, which
// includes cards with values from 0 to 12.
func GetNumberCards() []*Card {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	cards := []*Card{}

	for _, value := range numbers {
		cards = append(cards, &Card{
			Value:    value,
			IsNumber: true,
		})
	}

	return cards
}

// GetSpecialCards returns a slice of special cards, which includes
// Freeze, Flip 3 and Second Chance cards.
func GetSpecialCards() []*Card {
	cards := []*Card{}
	names := []string{"Freeze", "Flip 3", "Second Chance"}

	for i := range names {
		for range 3 {
			cards = append(cards, &Card{
				Value:    0,
				Category: names[i],
			})
		}
	}

	return cards
}
