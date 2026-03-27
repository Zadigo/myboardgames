package cards

type Card struct {
	// The card's value, for number cards this is the number, for multiplier
	// cards this is the multiplier, for bonus cards this is the bonus points,
	// for special cards this is the special effect
	value int
	// The player who owns the card, 0 for unowned
	owner int
	// nil, Freeze, Flip 3 or Second Chance
	category string
	// Is x2 instead of +2
	isMultiplier bool
	// Base card
	isNumber bool
	// One of x2, +2, +4, +6, +8 or +10
	isBonus bool
	// One of: Freeze, Flip 3 or Second Chance
	isSpecial bool
}

type Player struct {
	// Player's username
	username string
	// Unique identifier for the player
	uuid string
	// The table on which the player is currently playing
	tableUuid string
	// The current number of cards the player has flipped
	numberOfCards int
	// Whether the player was stopped by the freeze card
	isFreezed bool
	// The number of rounds played by the player
	numberOfRounds int
	// The cards that the player has hands
	cards []Card
}

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
