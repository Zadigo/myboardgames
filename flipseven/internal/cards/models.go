package cards

type Card struct {
	// The card's value, for number cards this is the number, for multiplier
	// cards this is the multiplier, for bonus cards this is the bonus points,
	// for special cards this is the special effect
	Value int `json:"value"`
	// The player who owns the card, 0 for unowned
	Owner int `json:"owner"`
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
