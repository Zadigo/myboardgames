package backend

func (c *MarvelCard) Buy() {
	// Implement the logic for buying the card
}

func (c *MarvelCard) Reserve() {
	// Implement the logic for reserving the card
}

func MarvelCardsLevelOne() []CardInterface {
	characters := []map[string]any{
		{"name": "LockJaw", "points": 0, "mind": 1, "space": 1, "soul": 1, "power": 1, "reality": 1, "time": 1, "shield": 1},
	}
	return CreateCard(false, 1, characters)
}

func MarvelCardsLevelTwo() []CardInterface {
	characters := []map[string]any{
		{"name": "Miles Morales", "points": 0, "mind": 1, "space": 1, "soul": 1, "power": 1, "reality": 1, "time": 1, "shield": 1},
	}
	return CreateCard(false, 2, characters)
}

func MarvelCardsLevelThree() []CardInterface {
	characters := []map[string]any{
		{"name": "Luke Cage", "points": 5, "mind": 1, "space": 1, "soul": 1, "power": 1, "reality": 1, "time": 1, "shield": 1},
	}
	return CreateCard(false, 3, characters)
}
