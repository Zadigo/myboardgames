package games

import "github.com/Zadigo/basegame/internal/models"

type GameInterface interface {
	GetUuid() string
	// Start initializes and starts the game server application.
	// It checks if the game instance is properly initialized, prepares the game state,
	// and then starts the game logic.
	Start() error
	Stop() error
	// AddPlayer adds a new player to the game. It takes a WebsocketClient representing the player
	// CreateCards initializes the card deck for the game,
	// setting up the necessary cards and their properties. This method is
	// responsible for ensuring that the game has a complete set of cards ready for play,
	// and it may involve shuffling or organizing the cards as needed.
	CreateCards() error
	// Prepare sets up the game state, initializes necessary components,
	// and ensures that the game is ready to start. This method is called before
	// the game begins and can include tasks such as shuffling cards, assigning
	// players to tables, and setting initial game parameters.
	Prepare() error
	// AddPlayer adds a new player to the game. It takes a WebsocketClient representing the player
	// and adds them to the game's player list. This method is responsible for managing
	// player connections and ensuring that new players can join the game seamlessly.
	AddPlayer(player *WebsocketClient) error
	// RemovePlayer removes a player from the game. It takes a WebsocketClient representing the player
	// and removes them from the game's player list. This method is responsible for managing
	// player disconnections and ensuring that players can leave the game gracefully.
	RemovePlayer(player *WebsocketClient) error
	// GetPlayer retrieves a player from the game based on their unique identifier (UUID).
	// It takes a playerUuid string and returns the corresponding WebsocketClient if found,
	// or an error if the player does not exist in the game. This method is useful for
	// managing player interactions and accessing player-specific information during gameplay.
	GetPlayer(playerUuid string) (*WebsocketClient, error)
	GetCurrentPlayer() (*WebsocketClient, error)
	SetCurrentPlayer(clientUuid string) error
	NextRound()
	NotifyAll()
	// SetOptions sets the game options for the game instance. It takes a GameAppOptions struct
	// containing configuration settings and parameters for the game. This method allows
	// for customization of game behavior, rules, and other settings based on the provided options.
	SetOptions(options models.GameAppOptions)
	// GetOptions retrieves the current game options for the game instance.
	// It returns a GameAppOptions struct containing the configuration settings and parameters
	// that are currently applied to the game. This method allows for inspection of the game's
	// configuration and can be used to adjust gameplay or settings as needed.
	GetOptions() models.GameAppOptions
}
