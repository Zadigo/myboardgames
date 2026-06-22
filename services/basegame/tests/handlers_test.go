package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuccessHandlerSuite struct {
	suite.Suite
}

func (suite *SuccessHandlerSuite) SetupTest() {}

func (suite *SuccessHandlerSuite) TestCreateGameHandler() {
	server, recorder := NewCreateRecorder(suite.T())
	defer server.Close()

	suite.NotNil(server)
	suite.NotNil(recorder)
}

func (suite *SuccessHandlerSuite) TestJoinGameHandler() {
	conn, server, recorder, err := NewJoinRecorder(suite.T())
	defer server.Close()
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.NotNil(conn)
	suite.NotNil(server)
	suite.NotNil(recorder)
}

func (suite *SuccessHandlerSuite) TestObserveGameHandler() {
	conn, server, recorder, err := NewObserveRecorder(suite.T())
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
