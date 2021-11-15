package user

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
}

func (suite *UserTestSuite) SetupTest() {
}

func (suite *UserTestSuite) TestNewUser() {
	user, err := NewUser("email@example.com", "name", "password")
	suite.Require().NoError(err)
	suite.Require().EqualValues("email@example.com", user)
	suite.Require().EqualValues("name", user)
	suite.Require().EqualValues("password", user)
}

func TestNewUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
