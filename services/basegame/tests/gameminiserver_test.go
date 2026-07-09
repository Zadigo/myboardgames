package tests

import (
	"context"
	"testing"

	"github.com/Zadigo/basegame/internal/game"
	"github.com/Zadigo/basegame/internal/game/games"
	"github.com/Zadigo/basegame/internal/models"
	"github.com/stretchr/testify/suite"
)

type TestGameMiniServer struct {
	suite.Suite
	gameApp *game.GamesMiniServer
}

func (suite *TestGameMiniServer) SetupTest() {
	ctx := context.WithValue(context.Background(), "baseDir", ".")

	suite.gameApp = game.NewGamesMiniServer(ctx, models.GameAppOptions{
		GameType:   games.STANDARD,
		AppOptions: models.AppOptions{},
		Debug:      true,
	})
}

func TestGamesMiniServer(t *testing.T) {
	suite.Run(t, new(TestGameMiniServer))
}
