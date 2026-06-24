package e2e

import (
	"fmt"
	"testing"

	"github.com/Zadigo/basegame/internal/handlers"
	"github.com/Zadigo/basegame/internal/models"
	"github.com/Zadigo/basegame/tests/utils"
	"github.com/stretchr/testify/suite"
)

type E2ETestSuite struct {
	suite.Suite
}

func (s *E2ETestSuite) SetupSuite() {}

func (s *E2ETestSuite) TearDownSuite() {}

func (s *E2ETestSuite) TestCompleteRoute() {
	gameId := "some-game-id"

	s.T().Run("Create game", func(t *testing.T) {})

	handler := handlers.GenericHandler{
		BaseHandler: handlers.BaseHandler{},
	}

	message := &models.WebsocketMessage{}
	conn, _, _, _ := utils.NewWebsocketRecorder(s.T(), fmt.Sprintf("/%s/join", gameId), handler.JoinGameHandler)
	s.NotNil(conn, "WebSocket connection should not be nil")

	// Receive the connection ID
	err := conn.ReadJSON(message)
	s.NoError(err, "Should not error when reading JSON")

	s.T().Run("Must identify", func(t *testing.T) {
		conn.WriteJSON(models.WebsocketMessage{
			BaseWebsocketMessage: models.BaseWebsocketMessage{
				Action: models.MUST_IDENTIFY,
			},
		})

		s.T().Run("Join", func(t *testing.T) {
			conn.WriteJSON(message)

			err = conn.ReadJSON(message)
			s.NoError(err, "Should not error when reading JSON")

			s.T().Run("Start game", func(t *testing.T) {})

			s.T().Run("Start turn", func(t *testing.T) {})

			s.T().Run("Flip card", func(t *testing.T) {})

			s.T().Run("End game", func(t *testing.T) {})
		})
	})
}

func (s *E2ETestSuite) TestWebSocketConnection() {
	suite.Run(s.T(), &E2ETestSuite{})
}
