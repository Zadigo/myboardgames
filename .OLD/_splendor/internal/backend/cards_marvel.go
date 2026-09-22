package backend

func (c *MarvelCard) Buy() {
	// Implement the logic for buying the card
}

func (c *MarvelCard) Reserve() {
	// Implement the logic for reserving the card
}

func MarvelCardsLevelOne() []CardInterface {
	characters := []map[string]any{
		{"name": "LockJaw", "avengerTag": 0, "bonus": "space", "points": 0, "mind": 1, "space": 1, "soul": 1, "power": 1, "reality": 1, "time": 1, "shield": 0},
		{"name": "Kate Bishop", "avengerTag": 1, "bonus": "soul", "points": 0, "mind": 0, "space": 1, "soul": 0, "power": 1, "reality": 1, "time": 0, "shield": 0},
		{"name": "Squirrel Girl", "avengerTag": 1, "bonus": "mind", "points": 0, "mind": 0, "space": 0, "soul": 1, "power": 2, "reality": 2, "time": 0, "shield": 0},
		{"name": "Kate Bishop", "avengerTag": 0, "bonus": "soul", "points": 0, "mind": 0, "space": 0, "soul": 0, "power": 0, "reality": 0, "time": 0, "shield": 0},
	}
	return CreateCard(false, 1, characters)
}

func MarvelCardsLevelTwo() []CardInterface {
	characters := []map[string]any{
		{"name": "Miles Morales", "avengerTag": 0, "bonus": "mind", "points": 0, "mind": 1, "space": 1, "soul": 1, "power": 1, "reality": 1, "time": 1, "shield": 0},
	}
	return CreateCard(false, 2, characters)
}

func MarvelCardsLevelThree() []CardInterface {
	characters := []map[string]any{
		{"name": "Luke Cage", "avengerTag": 0, "bonus": "power", "points": 5, "mind": 1, "space": 1, "soul": 1, "power": 1, "reality": 1, "time": 1, "shield": 0},
	}
	return CreateCard(false, 3, characters)
}
