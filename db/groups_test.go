package db

import (
	"github.com/kamva/mgm/v3"
	"github.com/stretchr/testify/suite"
	"github.com/xBlaz3kx/userManagementExample/internal/configuration"
	"github.com/xBlaz3kx/userManagementExample/model/group"
	"github.com/xBlaz3kx/userManagementExample/model/user"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

type groupDatabaseTestSuite struct {
	suite.Suite
	groupId string
	userId  string
}

func (suite *groupDatabaseTestSuite) SetupTest() {
	mgm.Coll(&group.Group{}).DeleteMany(mgm.Ctx(), bson.M{})
	mgm.Coll(&user.User{}).DeleteMany(mgm.Ctx(), bson.M{})

	usr, err := AddUser(user.User{
		Email:    "examplemail@example.com",
		Password: "examplePass",
		Name:     "exampleUser",
	})
	suite.Assert().NoError(err)
	suite.userId = usr.ID.Hex()

	addedGroup, err := AddGroup(group.Group{
		Name:  "exampleGroup",
		Users: []string{suite.userId},
	})

	suite.Assert().NoError(err)
	suite.groupId = addedGroup.ID.Hex()
}

func (suite *groupDatabaseTestSuite) TestAddGroup() {
	newGroup, err := group.NewGroup("exampleGroup1")
	suite.Assert().NoError(err)

	_, err = AddGroup(*newGroup)
	suite.Assert().NoError(err)
}

func (suite *groupDatabaseTestSuite) TestGetGroup() {
	// group initialized in the beginning of the test
	exampleGroup, err := GetGroup(suite.groupId)
	suite.Assert().NoError(err)
	suite.Assert().NotNil(exampleGroup)
	suite.Assert().EqualValues("exampleGroup", exampleGroup.Name)
	suite.Assert().EqualValues(suite.groupId, exampleGroup.ID.Hex())

	// no group with the name
	exampleGroup, err = GetGroup("exampleGroup123")
	suite.Assert().Error(err)
}

func (suite *groupDatabaseTestSuite) TestGetGroups() {
	groups, err := GetGroups()
	suite.Assert().NoError(err)
	suite.Assert().NotNil(groups)
	suite.Assert().EqualValues(suite.groupId, groups[0].ID.Hex())
}

func (suite *groupDatabaseTestSuite) TestUpdateGroup() {
	newName := "newGroupName"
	dGroup, err := UpdateGroup(suite.groupId, &newName)
	suite.Assert().NoError(err)

	// check if the group is updated
	updatedGroup, err := GetGroup(suite.groupId)
	suite.Assert().NoError(err)
	suite.Assert().NotNil(updatedGroup)
	suite.Assert().EqualValues(dGroup.ID.Hex(), updatedGroup.ID.Hex())

	//update with nil name
	updatedGroup, err = UpdateGroup(suite.groupId, nil)
	suite.Assert().NoError(err)

	// group should have the same name
	updatedGroup, err = GetGroup(suite.groupId)
	suite.Assert().NoError(err)
	suite.Assert().NotNil(updatedGroup)
	suite.Assert().EqualValues(suite.groupId, updatedGroup.ID.Hex())
}

func (suite *groupDatabaseTestSuite) TestDeleteGroup() {
	deleted, err := DeleteGroup(suite.groupId)
	suite.Assert().NoError(err)
	suite.Assert().True(deleted)

	deleted, err = DeleteGroup("nonExistingGroupId")
	suite.Assert().NoError(err)
	suite.Assert().False(deleted)
}

func (suite *groupDatabaseTestSuite) TestAddUserToGroup() {
	grp, err := AddUserToGroup(suite.groupId, "exampleUser12")
	suite.Assert().NoError(err)
	suite.Assert().NotNil(grp)

	grp, err = AddUserToGroup(suite.groupId, suite.userId)
	suite.Assert().Error(err)
}

func (suite *groupDatabaseTestSuite) TestRemoveUserFromGroup() {
	grp, err := RemoveUserFromGroup(suite.groupId, suite.userId)
	suite.Assert().NoError(err)
	suite.Assert().NotNil(grp)

	grp, err = RemoveUserFromGroup(suite.groupId, "exampleUser123")
	suite.Assert().Error(err)
}

func TestGroups(t *testing.T) {
	Connect(configuration.Mongo{
		Host:     "localhost",
		Username: "root",
		Password: "examplepass",
		Port:     27017,
		Database: "userManagement",
	})

	suite.Run(t, new(groupDatabaseTestSuite))
}
