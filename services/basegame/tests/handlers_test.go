package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuccessHandlerSuite struct {
	suite.Suite
}

func (suite *SuccessHandlerSuite) SetupTest() {}

func (suite *SuccessHandlerSuite) TestCreateGameHandler()  {}
func (suite *SuccessHandlerSuite) TestJoinGameHandler()    {}
func (suite *SuccessHandlerSuite) TestObserveGameHandler() {}

func TestSuccessHandlerSuite(t *testing.T) {
	suite.Run(t, new(SuccessHandlerSuite))
}
