package tests

import (
	"context"
	"testing"
	"time"

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

func (suite *TestGameMiniServer) TestCreateAndRetrieveGame() {
	suite.Run("should create a new game", func() {
		game := suite.gameApp.CreateGame(models.GameAppOptions{
			GameType: games.STANDARD,
		})

		suite.NotNil(game)
		suite.Len(suite.gameApp.Games, 1)

		suite.Run("should be able to get game", func() {
			result, err := suite.gameApp.GetGame(game.GetUuid())
			suite.NotNil(result)
			suite.NoError(err)
			suite.Equal(game.GetUuid(), result.GetUuid())
		})
	})
}

func (suite *TestGameMiniServer) TestStartAndStopServer() {
	// ctx, cancel := context.WithDeadline(suite.T().Context(), time.Now().Add(10*time.Second))
	parentCtx := suite.T().Context()
	parentCtx = context.WithValue(parentCtx, "gameType", games.STANDARD)
	// ctx, cancel := context.WithTimeout(parentCtx, 15*time.Second)

	suite.gameApp.SetContext(parentCtx)

	suite.Run("should start/stop the server", func() {
		// defer cancel()

		suite.gameApp.Start()
		suite.True(suite.gameApp.IsRunning.Load())

		time.Sleep(2 * time.Second) // Simulate some server activity
	})

	//	suite.Run("should stop the server", func() {
	//		suite.gameApp.Stop()
	//		suite.False(suite.gameApp.IsRunning.Load())
	//	})
}

func TestGamesMiniServer(t *testing.T) {
	suite.Run(t, new(TestGameMiniServer))
}
