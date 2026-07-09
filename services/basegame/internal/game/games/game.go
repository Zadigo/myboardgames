package games

import (
	"context"

	"github.com/Zadigo/basegame/internal/models"
)

const (
	STANDARD  = "standard"
	EXTENSION = "extension"
)

// CreateGame is a factory function that creates a new game instance based on the specified game type.
// It takes a context.Context and GameAppOptions as parameters and returns an instance of GameInterface,
// which can be either a StandardGame or an ExtensionGame, depending on the provided gameType.
// If the gameType is not recognized, it returns nil.
func CreateGame(ctx context.Context, options models.GameAppOptions) GameInterface {
	gameCtx := context.WithValue(ctx, "gameType", options.GameType)

	var game GameInterface

	switch options.GameType {
	case STANDARD:
		game = &StandardGame{
			BaseGame: BaseGame{
				ctx:     gameCtx,
				players: make(map[string]*WebsocketClient),
			},
		}
	case EXTENSION:
		game = &ExtensionGame{
			BaseGame: BaseGame{
				ctx:     gameCtx,
				players: make(map[string]*WebsocketClient),
			},
		}
	default:
		return nil
	}

	game.SetOptions(options)
	return game
}
