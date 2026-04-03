package models

type Card struct {
	Emerald  int
	Diamond  int
	Sapphire int
	Onyx     int
	Ruby     int
	Points   int
	Owner    *Player
}

func CardsLevelThree() []*Card {
	return []*Card{
		{Points: 0, Emerald: 1, Diamond: 1, Sapphire: 1, Onyx: 1, Ruby: 1},
	}
}

func CardsLevelTwo() []*Card {
	return []*Card{
		{Points: 0, Emerald: 1, Diamond: 1, Sapphire: 1, Onyx: 1, Ruby: 1},
	}
}

func CardsLevelOne() []*Card {
	return []*Card{
		{Points: 0, Emerald: 1, Diamond: 1, Sapphire: 1, Onyx: 1, Ruby: 1},
	}
}
