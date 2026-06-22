package game

const (
	STANDARD  = "standard"
	EXTENSION = "extension"
)

type GameApp struct {
	game GameInterface
}

func (app *GameApp) Start() error {
	if app.game == nil {
		return nil
	}
	app.game.Prepare()
	
	return app.game.Start()
}

func NewGameApp(gameType string) *GameApp {
	game := CreateGame(gameType)
	return &GameApp{game: game}
}
