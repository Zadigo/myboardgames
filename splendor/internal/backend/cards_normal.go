package backend

func (c *NormalCard) Buy() {
	// Implement the logic for buying the card
}

func (c *NormalCard) Reserve() {
	// Implement the logic for reserving the card
}

func NormalCardsLevelOne() []CardInterface {
	characters := []map[string]any{
		{"name": "", "points": 0, "emerald": 1, "diamond": 1, "sapphire": 1, "onyx": 1, "ruby": 1},
	}
	return CreateCard(true, 1, characters)
}

func NormalCardsLevelTwo() []CardInterface {
	characters := []map[string]any{
		{"name": "", "points": 0, "emerald": 1, "diamond": 1, "sapphire": 1, "onyx": 1, "ruby": 1},
	}
	return CreateCard(true, 2, characters)
}

func NormalCardsLevelThree() []CardInterface {
	characters := []map[string]any{
		{"name": "", "points": 5, "emerald": 1, "diamond": 1, "sapphire": 1, "onyx": 1, "ruby": 1},
	}
	return CreateCard(true, 3, characters)
}
