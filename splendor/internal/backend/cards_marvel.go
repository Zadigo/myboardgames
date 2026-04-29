package backend

func (c *MarvelCard) Buy() {
	// Implement the logic for buying the card
}

func (c *MarvelCard) Reserve() {
	// Implement the logic for reserving the card
}

func CardsLevelOne() []CardInterface {
	characters := []map[string]any{
		{"name": "LockJaw", "points": 0, "emerald": 1, "diamond": 1, "sapphire": 1, "onyx": 1, "ruby": 1},
	}
	return CreateCard(MarvelCard{}, 1, characters)
}

func CardsLevelTwo() []CardInterface {
	characters := []map[string]any{
		{"name": "Miles Morales", "points": 0, "emerald": 1, "diamond": 1, "sapphire": 1, "onyx": 1, "ruby": 1},
	}
	return CreateCard(MarvelCard{}, 2, characters)
}

func CardsLevelThree() []CardInterface {
	characters := []map[string]any{
		{"name": "Luke Cage", "points": 5, "emerald": 1, "diamond": 1, "sapphire": 1, "onyx": 1, "ruby": 1},
	}
	return CreateCard(MarvelCard{}, 3, characters)
}
