package models

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type AppConfigInterface interface {
	LoadConfig(app AppInterface) error
}

type AppExtraInterface interface {
	GetRedisClient() *redis.Client
	GetBaseDir() string
	GetServerApp() ServerAppInterface
}

// Deprecated: Do we really need this interface?
type AppInterface interface {
	AppExtraInterface
	Start() error
	GetContext() context.Context
}

type ServerAppInterface interface {
	// JoinGame allows a player to join a specific game instance.
	JoinGame(conn *websocket.Conn, gameType string, gameUUID string) error
	// CreateGame allows the creation of a new game instance of a specific type.
	CreateGame(gameType string) (newGame GameInterface, err error)
	// GetMiniServer(gameType string) *game.GamesMiniServer
	// GetAllMiniServers() map[string]*game.GamesMiniServer
}

type PlayerInterface interface {
	GetCards() []CardInterface
	SetCards(cards []CardInterface)
	GetTotalScore() int
	GetIsFrozen() bool
	SetIsFrozen(frozen bool)
	GetNumberOfCards() int
}

type WebsocketClientInterface interface {
	GetConn() *websocket.Conn
	GetPlayer() PlayerInterface
	GetUuid() string
	SendJsonMessage(message WebsocketMessage) error
	ReceiveJsonMessage() (WebsocketMessage, error)
	GetMessagesChannel() chan WebsocketMessage
	CloseMessagesChannel()
	SendMessageToChannel(message WebsocketMessage)
}

type GameInterface interface {
	GetUuid() string
	// Start initializes and starts the game server application.
	// It checks if the game instance is properly initialized, prepares the game state,
	// and then starts the game logic.
	Start() error
	// Stop terminates the game instance, performing any necessary cleanup and ensuring that the game is properly closed.
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
	AddPlayer(player WebsocketClientInterface) error
	// RemovePlayer removes a player from the game. It takes a WebsocketClient representing the player
	// and removes them from the game's player list. This method is responsible for managing
	// player disconnections and ensuring that players can leave the game gracefully.
	RemovePlayer(player WebsocketClientInterface) error
	// GetPlayer retrieves a player from the game based on their unique identifier (UUID).
	// It takes a playerUuid string and returns the corresponding WebsocketClient if found,
	// or an error if the player does not exist in the game. This method is useful for
	// managing player interactions and accessing player-specific information during gameplay.
	GetPlayer(playerUuid string) (WebsocketClientInterface, error)
	// GetCurrentPlayer retrieves the player whose turn it is currently in the game.
	GetCurrentPlayer() (WebsocketClientInterface, error)
	// SetCurrentPlayer sets the player whose turn it is currently in the game.
	SetCurrentPlayer(clientUuid string) error
	// NextRound advances the game to the next round, updating the game state accordingly.
	NextRound()
	// NotifyAll sends a notification to all players in the game, typically used to broadcast game events or updates.
	NotifyAll()
	// SetOptions sets the game options for the game instance. It takes a GameAppOptions struct
	// containing configuration settings and parameters for the game. This method allows
	// for customization of game behavior, rules, and other settings based on the provided options.
	SetOptions(options GameAppOptions)
	// GetOptions retrieves the current game options for the game instance.
	// It returns a GameAppOptions struct containing the configuration settings and parameters
	// that are currently applied to the game. This method allows for inspection of the game's
	// configuration and can be used to adjust gameplay or settings as needed.
	GetOptions() GameAppOptions
}

// Additional options that may be needed
// for resolving card effects in the future
// useful for special cards that may have
// more complex effects
type CardResolveOptions struct {
	// Uuid of the player to whom the
	// card effect is being applied, if applicable
	ToPlayer string
}

type CardInterface interface {
	GetUuid() string
	// Resolution logic for the card's effect, if any
	Resolve(currentScore int, options CardResolveOptions) int
	// Set the player UUID to indicate ownership of the card
	SetPlayer(playerUuid string)
	// Get the player UUID of the player who owns the card
	GetPlayer() string
	// Get the value of the card
	GetValue() int
	// Check if the card is a special card (e.g., freeze, flip3, second_chance)
	IsSpecial() bool
	// Get the operator of the card (e.g., "+", "-", "*", "/", "freeze", "flip3", "second_chance")
	GetOperator() string
	// Check card operation type (e.g., "draw", "flip3", "freeze", "second_chance")
	IsCardOperation(operator string) bool
}
