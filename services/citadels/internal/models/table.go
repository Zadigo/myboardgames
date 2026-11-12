package models

import (
	"cmp"
	"fmt"
	"slices"
	"sync"
	"sync/atomic"
)

func Filter[T ~[]E, E any](values T, f func(E) bool) []E {
	items := []E{}
	for _, v := range values {
		result := f(v)
		if result {
			items = append(items, v)
		}
	}
	return items
}

type PileInterface[T any] interface {
	Shuffle()
	GetFirstCard() T
	GetXCards(k int) []T
	PlaceCardAtBottom(cards ...T)
	All() []T
	SelectedCards() []T
	NoneSelectedCards() []T
	HasCards() bool
	Build()
}

type BaseTable struct {
	// A helper to track character resolution order. There can
	// only be 8 characters at most. With two players, the
	// Emperor cannot be selected
	ResolutionIndicator []BaseCharacterIndicator
	ResolutionIndex     atomic.Int32
	charactersCards     []CharacterInterface
	Players             []PlayerInterface
	Characters          []string
	Crown               PlayerInterface
	Stage               string
	DistrictPile        PileInterface[DistrictInterface]
	Phase               string
	IsSetup             atomic.Bool
	mu                  sync.Mutex
}

func (b *BaseTable) GetCarachterCards() []CharacterInterface {
	return b.charactersCards
}

func (b *BaseTable) GetCharacterIndicators() []BaseCharacterIndicator {
	return b.ResolutionIndicator
}

func (b *BaseTable) GetState() bool {
	return b.IsSetup.Load()
}

func (b *BaseTable) GetCurrentCharacterHand() []CharacterInterface {
	return Filter(b.charactersCards, func(c CharacterInterface) bool {
		return c.GetPlayer() != nil
	})
}

// func (b *BaseTable) GetNoneSelectedDistricts() []DistrictInterface {
// 	pile := b.DistrictPile.All()
// 	index := slices.IndexFunc(pile, func(c DistrictInterface) bool {
// 		return c.GetPlayer() != nil
// 	})
// 	return pile[index:]
// }

// func (b *BaseTable) GetSelectedDistricts() []DistrictInterface {
// 	pile := b.DistrictPile.All()
// 	index := slices.IndexFunc(pile, func(c DistrictInterface) bool {
// 		return c.GetPlayer() != nil
// 	})
// 	return pile[:index]
// }

func (b *BaseTable) AddPlayer(player PlayerInterface) (remainingPlaces int) {
	b.Players = append(b.Players, player)
	return 8 - len(b.Players)
}

// CreateCharacterIndicators creates the memo tokens that will be used to remember the
// current character that is being resolved
func (b *BaseTable) CreateCharacterIndicators(characters []CharacterInterface) error {
	if len(characters) > 8 {
		return fmt.Errorf("Can only have 8 characters at most")
	}

	for i, n := range characters {
		b.ResolutionIndicator[i] = BaseCharacterIndicator{
			Name:  n.GetName(),
			Level: n.GetLevel(),
		}
	}

	b.charactersCards = characters

	slices.SortFunc(b.ResolutionIndicator, func(a, b BaseCharacterIndicator) int {
		return cmp.Compare(a.Level, b.Level)
	})

	return nil
}

func (b *BaseTable) DistributeDistricts(k int) {
	if !b.DistrictPile.HasCards() {
		b.DistrictPile.Build()
	}

	for _, p := range b.Players {
		cards := b.DistrictPile.GetXCards(4)
		p.DrawDistricts(cards)
	}
}

func (b *BaseTable) GiveCrown(mode string) error {
	if len(b.Players) == 0 {
		return fmt.Errorf("no players")
	}

	switch mode {
	case "random":
		// Select a random player
		return nil
	case "oldest_player":
		// Selet the oldest player
		return nil
	default:
		return fmt.Errorf("Could not recnognize mode")
	}
}

func (b *BaseTable) GetPlayerWithCrown() PlayerInterface {
	for _, p := range b.Players {
		if p.HasCrown() {
			return p
		}
	}

	return nil
}

// StartRound
func (b *BaseTable) StartRound() {
	b.IsSetup.Store(true)
}

func (b *BaseTable) Setup(characters []CharacterInterface, mode string) (state bool, currentPlayer PlayerInterface) {
	b.CreateCharacterIndicators(characters)
	b.DistributeDistricts(4)
	b.GiveCrown(mode)
	b.StartRound()
	return b.IsSetup.Load(), b.GetPlayerWithCrown()
}
