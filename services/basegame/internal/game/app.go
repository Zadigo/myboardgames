package game

import "fmt"

const (
	STANDARD  = "standard"
	EXTENSION = "extension"
)

type GameApp struct {
	game GameInterface
}

func (app *GameApp) Start() error {
	if app.game == nil {
		return fmt.Errorf("❌ Game server is not initialized")
	}

	app.game.Prepare()
	app.game.Start()
	
	return nil
}

func NewGameApp(gameType string) *GameApp {
	game := CreateGame(gameType)
	return &GameApp{game: game}
}
