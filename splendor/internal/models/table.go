package models

import (
	"sync"
)

type PlayingTable struct {
	DeckLevelOne   []*Card
	DeckLevelTwo   []*Card
	DeckLevelThree []*Card

	CardsLevelThree []*Card
	CardsLevelTwo   []*Card
	CardsLevelOne   []*Card

	Emerald   int
	Diamond   int
	Sapphire  int
	Onyx      int
	Ruby      int
	GoldJoker int

	GameStarted bool
	Players     map[string]*Player
}

type LayerInterface interface {
	AddCard(card *Card, level int)
}

type TableLayer struct {
	Layer          LayerInterface
	CurrentRound   int
	CurrentPlayer  *Player
	MaxPlayers     int
	MaxCardsPerRow int
	mu             sync.Mutex
}

// Add a player to the table.
func (t *PlayingTable) AddPlayer(player *Player) {
	if t.GameStarted {
		t.Players[player.Uuid] = player
	}
}

// Start the game and initialize the table.
func (t *PlayingTable) StartGame() {
	if !t.GameStarted && t.NumberOfPlayers() > 2 {
		t.GameStarted = true

		if t.NumberOfPlayers() == 2 {
			t.Emerald = 4
			t.Diamond = 4
			t.Sapphire = 4
			t.Onyx = 4
			t.Ruby = 4
		}

		if t.NumberOfPlayers() == 3 {
			t.Emerald = 5
			t.Diamond = 5
			t.Sapphire = 5
			t.Onyx = 5
			t.Ruby = 5
		}

		if t.NumberOfPlayers() == 4 {
			t.Emerald = 7
			t.Diamond = 7
			t.Sapphire = 7
			t.Onyx = 7
			t.Ruby = 7
		}
	}
}

// End the game and reset the table.
func (t *PlayingTable) EndGame() {
	if t.GameStarted {
		t.GameStarted = false
		t.Players = make(map[string]*Player)

		t.CardsLevelThree = []*Card{}
		t.CardsLevelTwo = []*Card{}
		t.CardsLevelOne = []*Card{}

		t.Emerald = 7
		t.Diamond = 7
		t.Sapphire = 7
		t.Onyx = 7
		t.Ruby = 7
		t.GoldJoker = 5
	}
}

func (t *PlayingTable) AddCard(card *Card, level int) {

}

func (t *PlayingTable) NewDeck() {
	cardsLevelOne := CardsLevelOne()
	cardsLevelTwo := CardsLevelTwo()
	cardsLevelThree := CardsLevelThree()

	t.DeckLevelOne = append(t.DeckLevelOne, cardsLevelOne...)
	t.DeckLevelTwo = append(t.DeckLevelTwo, cardsLevelTwo...)
	t.DeckLevelThree = append(t.DeckLevelThree, cardsLevelThree...)
}

// Return the number of players currently at the table.
func (t *PlayingTable) NumberOfPlayers() int {
	return len(t.Players)
}

// Create the initial influnce queue with 5 cards of
// each color for each player
func CreateTableLayer() *TableLayer {
	return &TableLayer{
		CurrentRound:   0,
		CurrentPlayer:  nil,
		MaxPlayers:     4,
		MaxCardsPerRow: 4,
		Layer: &PlayingTable{
			CardsLevelThree: []*Card{},
			CardsLevelTwo:   []*Card{},
			CardsLevelOne:   []*Card{},
		},
	}
}
