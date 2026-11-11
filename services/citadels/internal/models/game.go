package models

import (
	"sync"
)

const (
	SelectionPhase string = "selection_phase"
	TurnPhase      string = "turn_phase"
)

// This section represents all the phases that occur when the websocket
// communicates with Nuxt. Eeach action/phase allows the websocket to
// implement a certain logic
const (
	WaitingForPlayers string = "waiting_for_players"
	// Stage used when initializing the game
	// 	1. Place the character reminder cards on the table in order of rank
	//	2. Shuffle the 14 unique district cards with the 54 basic ones
	//  3. Deal 4 cards to each player
	//  4. Place the remaining district cards in a facedown pile on the table
	//  5. Create the bank with all the gold coins
	//  6. Give each player 2 golds from the bank (it's their starting gold)
	//  7. The oldest character takes the crown and the eight character cards for the round
	ApplyCharacterSelection string = "apply_character_selection"
	CreateDistricts         string = "create_districts"
	DetermineCrown          string = "determine_crown"
	DistributeFourDistricts string = "distribute_four_districts"
	StartRound              string = "start_round"
	GiveCharacterCards      string = "give_character_cards"
	// Stage to be used whenever the game is in a
	// waiting state e.g. waiting for players...
	WaitingForPlayerChoice string = "waiting_for_player_choice"
	// Stage used when the game discards characters face up
	// and/or one character face down before having the players select
	// their character for the current round
	// 	- 4 Players: 2 (faceup), 1 (facedown)
	//  - 5 Players: 1 (faceup), 1 (facedown)
	//  - 6 Players: 0 (faceup), 1 (facedown)
	//  - 7 Players: 0 (faceup), 1 (facedown)
	//
	// The King, Emporor and Patrician CANNOT be among the
	DiscardCharacters string = "discard_characters"
	// Stage used when the players are selecting their characters
	SelectCharacter string = "select_character"
	NextPlayer      string = "next_player"
	// The computer goes up the ladder one by one to check who has the current character
	SequentialCharacterResolution string = "sequential_character_resolution"
	// The computer moves up the ladder by one to resolve a character
	CharacterResolve string = "character_resolve"
	// Choice made by the player on his turn: two golds, two districts or build
	PlayerChoiceOnTurn string = "player_choice_on_turn"
	CheckGameState     string = "check_game_state"
	NextRound          string = "next_round"
	EndRound           string = "end_round"
	// Stage used one a player has completed 7 or more districts. The game
	// ends after the current round is finished.
	//
	// 	- Score points equal to the building cost of each of your districts.
	// 	- If your city has at least one district of each type, score 3 points.
	// 	- The player who first completed their city scores 4 points
	// 	- Any other player who completed their city scores 2 points
	// 	- Score any extra points from your unique districts
	GameEnd string = "game_end"
)

type PileInterface[T any] interface {
	Shuffle()
	GetFirstCard() T
	GetXCards(k int) []T
	PlaceCardAtBottom(cards ...T)
}

type GameInterface interface {
	GetPlayers() []PlayerInterface
	GetStage() string
	GetCrownOwner() PlayerInterface
	GetCharacters() []string
	GetID() string
	Start() error
	GetTable() *BaseTable
}

type BaseTable struct {
	// A helper to track character resolution order. There can
	// only be 8 characters at most. With two players, the
	// Emperor cannot be selected
	ResolutionIndicator []BaseCharacterIndicator
	Players             []PlayerInterface
	Characters          []string
	Crown               PlayerInterface
	Stage               string
	DistrictPile        PileInterface[DistrictInterface]
	Phase               string
	mu                  sync.Mutex
}

func (b *BaseTable) AddPlayer(player PlayerInterface) (remainingPlaces int) {
	for i := range b.Players {
		if b.Players[i] == nil {
			b.Players[i] = player
			break
		}
	}

	for i := range b.Players {
		if b.Players[i] == nil {
			remainingPlaces++
		}
	}

	return remainingPlaces
}

// BaseGame represents an ongoing game.
// Each game has a table for players to sit at.
type BaseGame struct {
	Table *BaseTable
	ID    string
}

func (g *BaseGame) GetID() string {
	return g.ID
}

func (g *BaseGame) GetPlayers() []PlayerInterface {
	return g.Table.Players
}

func (g *BaseGame) GetStage() string {
	return g.Table.Stage
}

func (g *BaseGame) GetCrownOwner() PlayerInterface {
	return g.Table.Crown
}

func (g *BaseGame) GetCharacters() []string {
	return g.Table.Characters
}

func (g *BaseGame) GetTable() *BaseTable {
	return g.Table
}
