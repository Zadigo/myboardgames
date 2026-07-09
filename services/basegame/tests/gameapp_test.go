package tests

import (
	"context"
	"testing"

	"github.com/Zadigo/basegame/internal/game/cards"
	"github.com/Zadigo/basegame/internal/game/games"
	"github.com/Zadigo/basegame/internal/handlers"
	"github.com/Zadigo/basegame/internal/models"
	"github.com/Zadigo/basegame/tests/utils"
	"github.com/stretchr/testify/suite"
)

type GameAppTestSuite struct {
	suite.Suite
	ctx     context.Context
	gameApp *games.GameApp
}

func (suite *GameAppTestSuite) SetupTest() {
	suite.ctx = context.WithValue(context.Background(), "baseDir", ".")

	suite.gameApp = games.NewGame(suite.ctx, models.GameAppOptions{
		GameType:   games.STANDARD,
		AppOptions: models.AppOptions{},
		Debug:      true,
	})
}

func (suite *GameAppTestSuite) TestNewGame() {
	gameApp := games.NewGame(suite.ctx, models.GameAppOptions{
		GameType:   games.STANDARD,
		AppOptions: models.AppOptions{},
		Debug:      true,
	})

	suite.NotNil(gameApp)
	suite.False(gameApp.IsRunning.Load())

	// err := suite.gameApp.Start()
	// suite.NoError(err)
	// suite.True(gameApp.IsRunning.Load())

	// err = suite.gameApp.Stop()
	// suite.NoError(err)
	// suite.False(suite.gameApp.IsRunning.Load())
}

func (suite *GameAppTestSuite) TestCreateWithNil() {
	suite.NotNil(suite.gameApp)

	oldGame := suite.gameApp
	suite.gameApp = nil

	err := suite.gameApp.Start()
	suite.Error(err)
	suite.EqualError(err, "❌ Game server is not initialized")

	suite.T().Cleanup(func() {
		suite.gameApp = oldGame
	})
}

func (suite *GameAppTestSuite) TestGetCurrentPlayer() {
	client := utils.CreateClientPlayer(suite.T())

	// TODO: Create a method to add a player to the game app and set the current player UUID
	suite.gameApp.CurrentPlayerUuid = client.GetUuid()
	suite.gameApp.Players[client.GetUuid()] = client

	player, err := suite.gameApp.GetCurrentPlayer()
	suite.NotNil(player)
	suite.NoError(err)
	suite.Equal(client.GetUuid(), player.GetUuid())
}

func (suite *GameAppTestSuite) TestGetCurrentPlayerWithNoCurrentPlayer() {
	// Ensure that the CurrentPlayerUuid is empty
	suite.gameApp.CurrentPlayerUuid = ""

	player, err := suite.gameApp.GetCurrentPlayer()
	suite.Nil(player)
	suite.Error(err)
	suite.EqualError(err, "❌ Current player UUID is not set")
}

func (suite *GameAppTestSuite) TestDrawCardFromPile() {
	suite.gameApp.CurrentPlayerUuid = ""

	handler := handlers.GenericHandler{}
	c1, _, _, _ := utils.NewWebsocketRecorder(suite.T(), "/join", handler.JoinGameHandler)
	c2, _, _, _ := utils.NewWebsocketRecorder(suite.T(), "/join", handler.JoinGameHandler)

	p1 := games.NewWebsocketClient(c1)
	p2 := games.NewWebsocketClient(c2)

	suite.gameApp.AddPlayer(p1)
	suite.gameApp.AddPlayer(p2)

	suite.Empty(suite.gameApp.CurrentPlayerUuid)

	suite.gameApp.SetCurrentPlayer(p1.GetUuid())

	// Create the card pile
	factory := cards.MakeCards(cards.STANDARD_CARD)
	suite.gameApp.CardPile.AddCards(factory.CreateCards()...)

	// Draw from the pile -
	err := suite.gameApp.DrawCardFromPile()

	suite.NoError(err)
	suite.False(suite.gameApp.MustResolve)
	suite.Equal(1, len(p1.Player.GetCards()))
	suite.False(p1.Player.IsFrozen)

	suite.T().Cleanup(func() {
		c1.Close()
		c2.Close()

		suite.gameApp.CardPile.Clear()
		suite.gameApp.CurrentPlayerUuid = ""
	})
}

func (suite *GameAppTestSuite) TestDrawCardFromPileWithSpecialCard() {
	// Pulling a special card from the pile should set MustResolve to true 
	// and populate EventResolutionOptions. The client should then resolve
	// the event before the game can continue.

	card1 := &cards.BaseCard{
		Uuid:         "card-1",
		CardType:     cards.STANDARD_SPECIAL_CARD,
		CardValue:    0,
		CardOperator: cards.OPERATION_FREEZE,
	}

	card2 := &cards.BaseCard{
		Uuid:         "card-2",
		CardType:     cards.STANDARD_SPECIAL_CARD,
		CardValue:    0,
		CardOperator: cards.OPERATION_FLIP3,
	}

	suite.gameApp.CardPile.AddCards(card1, card2)

	handler := handlers.GenericHandler{}
	c1, _, _, _ := utils.NewWebsocketRecorder(suite.T(), "/join", handler.JoinGameHandler)

	p1 := games.NewWebsocketClient(c1)

	err := suite.gameApp.AddPlayer(p1)
	suite.NoError(err)

	err = suite.gameApp.SetCurrentPlayer(p1.GetUuid())
	suite.NoError(err)

	type testCase struct {
		name string
		card *cards.BaseCard
	}

	testCases := []testCase{
		{name: "Draw Freeze Card", card: card1},
		{name: "Draw Flip3 Card", card: card2},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			err = suite.gameApp.DrawCardFromPile()

			suite.NoError(err)
			suite.True(suite.gameApp.MustResolve)
			suite.Equal(tc.card, suite.gameApp.EventResolutionOptions.Card)
			suite.Equal(p1, suite.gameApp.EventResolutionOptions.Player)
		})
	}

	suite.T().Cleanup(func() {
		suite.gameApp.MustResolve = false
		suite.gameApp.EventResolutionOptions = games.EventResolutionOptions{}
		suite.gameApp.CardPile.Clear()
	})
}

func (suite *GameAppTestSuite) TestDrawCardFromPileWithNoCardsLeft() {}

func (suite *GameAppTestSuite) TestDrawCardWithSameValue() {
	handler := handlers.GenericHandler{}

	c1, _, _, _ := utils.NewWebsocketRecorder(suite.T(), "/join", handler.JoinGameHandler)
	p1 := games.NewWebsocketClient(c1)

	card1 := &cards.BaseCard{
		Uuid:         "card-1",
		CardType:     cards.STANDARD_CARD,
		CardValue:    5,
		CardOperator: cards.OPERATION_ADD,
	}

	// Simulate that the player already has a card with the same value in their hand
	p1.Player.Cards = append(p1.Player.Cards, card1)

	suite.gameApp.AddPlayer(p1)
	suite.gameApp.SetCurrentPlayer(p1.GetUuid())

	card2 := &cards.BaseCard{
		Uuid:         "card-2",
		CardType:     cards.STANDARD_CARD,
		CardValue:    5,
		CardOperator: cards.OPERATION_ADD,
	}

	suite.gameApp.CardPile.AddCards(card2)

	err := suite.gameApp.DrawCardFromPile()
	suite.NoError(err)
	suite.True(p1.Player.IsFrozen)
}

func TestGameApp(t *testing.T) {
	suite.Run(t, new(GameAppTestSuite))
}
