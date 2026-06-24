package tests

import (
	"testing"

	"github.com/Zadigo/basegame/tests/utils"
	"github.com/stretchr/testify/suite"
)

type SuccessHandlerSuite struct {
	suite.Suite
}

func (suite *SuccessHandlerSuite) SetupTest() {}

func (suite *SuccessHandlerSuite) TestCreateGameHandler() {
	server, recorder := utils.NewCreateRecorder(suite.T())
	defer server.Close()

	suite.NotNil(server)
	suite.NotNil(recorder)
}

func (suite *SuccessHandlerSuite) TestJoinGameHandler() {
	conn, server, recorder, err := utils.NewJoinRecorder(suite.T())
	defer server.Close()
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.NotNil(conn)
	suite.NotNil(server)
	suite.NotNil(recorder)
}

func (suite *SuccessHandlerSuite) TestObserveGameHandler() {
	conn, server, recorder, err := utils.NewObserveRecorder(suite.T())
	defer server.Close()
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.NotNil(conn)
	suite.NotNil(server)
	suite.NotNil(recorder)
}

func TestSuccessHandlerSuite(t *testing.T) {
	suite.Run(t, new(SuccessHandlerSuite))
}
