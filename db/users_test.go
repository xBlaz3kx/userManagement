package db

import (
	"github.com/kamva/mgm/v3"
	"github.com/stretchr/testify/suite"
	"github.com/xBlaz3kx/userManagementExample/internal/configuration"
	"github.com/xBlaz3kx/userManagementExample/model/user"
	"go.mongodb.org/mongo-driver/bson"

	"testing"
)

type userDatabaseTestSuite struct {
	suite.Suite
	userId string
}

func (suite *userDatabaseTestSuite) SetupTest() {
	mgm.Coll(&user.User{}).DeleteMany(mgm.Ctx(), bson.M{})

	addedUser, err := AddUser(user.User{
		Email:    "example@email.com",
		Password: "examplePass",
		Name:     "example",
	})

	suite.Assert().NoError(err)
	suite.userId = addedUser.ID.Hex()
}

func (suite *userDatabaseTestSuite) TestAddUser() {
	addedUser, err := AddUser(user.User{
		Email:    "example2@email.com",
		Password: "examplePass",
		Name:     "example",
	})
	suite.Assert().NoError(err)
	suite.Assert().NotNil(addedUser)

	addedUser, err = AddUser(user.User{
		Email:    "example@email.com",
		Password: "examplePass",
		Name:     "example",
	})
	suite.Assert().Error(err)
}

func (suite *userDatabaseTestSuite) TestGetUser() {
	usr, err := GetUser(suite.userId)
	suite.Assert().NoError(err)
	suite.Assert().NotNil(usr)

	usr, err = GetUser("nonExistingUserId")
	suite.Assert().Error(err)
}

func (suite *userDatabaseTestSuite) TestGetUsers() {
	usrs, err := GetUsers()
	suite.Assert().NoError(err)
	suite.Assert().NotNil(usrs)
}

func (suite *userDatabaseTestSuite) TestUpdateUser() {
	newName := "newName"
	newPass := "newPassword"

	usr, err := UpdateUser(suite.userId, &newName, nil)
	suite.Assert().NoError(err)
	suite.Assert().Equal(suite.userId, usr.ID.Hex())
	suite.Assert().Equal(newName, usr.Name)

	usr, err = UpdateUser(suite.userId, nil, &newPass)
	suite.Assert().NoError(err)
	suite.Assert().Equal(suite.userId, usr.ID.Hex())
	suite.Assert().Equal(newPass, usr.Password)

	usr, err = UpdateUser("nonExistingUserId", nil, nil)
	suite.Assert().Error(err)
}

func (suite *userDatabaseTestSuite) TestDeleteUser() {
	deleted, err := DeleteUser(suite.userId)
	suite.Assert().NoError(err)
	suite.Assert().True(deleted)

	deleted, err = DeleteUser("nonExistingUserId")
	suite.Assert().False(deleted)
	suite.Assert().NoError(err)
}

func TestUsers(t *testing.T) {
	Connect(configuration.Mongo{
		Host:     "localhost",
		Username: "root",
		Password: "examplepass",
		Port:     27017,
		Database: "userManagement",
	})

	suite.Run(t, new(userDatabaseTestSuite))
}
