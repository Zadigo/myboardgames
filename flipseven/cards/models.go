package cards

type Card struct {
	// The card's value, for number cards this is the number, for multiplier
	// cards this is the multiplier, for bonus cards this is the bonus points,
	// for special cards this is the special effect
	Value int
	// The player who owns the card, 0 for unowned
	Owner int
	// nil, Freeze, Flip 3 or Second Chance
	Category string
	// Is x2 instead of +2
	IsMultiplier bool
	// Base card
	IsNumber bool
	// One of x2, +2, +4, +6, +8 or +10
	IsBonus bool
	// One of: Freeze, Flip 3 or Second Chance
	IsSpecial bool
}
