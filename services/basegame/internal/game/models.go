package game

import (
	"github.com/Zadigo/basegame/internal/game/cards"
	"github.com/Zadigo/basegame/internal/game/games"
)

const (
	EVENT_RESOLVE_CARD = "resolve_card"
	EVENT_RESOLVE_PLAYERS_FROZEN = "resolve_players_frozen"
)

type EventResolutionOptions struct {
	// The card that triggered the event or action
	// that must be resolved
	Card cards.CardInterface
	// The player that triggered the event or action
	Player *games.Player
}

type EventResolutionAction struct {
	Event string
	// A player to whom an event or action is directed
	ToPlayer *games.Player
}
