package game

import "fmt"

const (
	STANDARD  = "standard"
	EXTENSION = "extension"
)

type GameApp struct {
	game GameInterface
}

// Start initializes and starts the game server application. 
// It checks if the game instance is properly initialized, prepares the game state, 
// and then starts the game logic.
func (app *GameApp) Start() error {
	if app.game == nil {
		return fmt.Errorf("❌ Game server is not initialized")
	}

	app.game.Prepare()
	app.game.Start()
	
	return nil
}

// NewGameApp creates a new instance of the game server application 
// based on the specified game type (standard or extension). This server
// is independent of the main chi server and can be used to manage game sessions, 
// handle player interactions, and maintain game state.
func NewGameApp(gameType string) *GameApp {
	game := CreateGame(gameType)
	return &GameApp{game: game}
}
