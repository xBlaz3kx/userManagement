package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"github.com/xBlaz3kx/userManagementExample/db"
	"github.com/xBlaz3kx/userManagementExample/internal/configuration"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type (
	Group struct {
		Id    string   `json:"id"`
		Name  string   `json:"name"`
		Users []string `json:"users"`
	}

	User struct {
		Id       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	apiTestSuite struct {
		suite.Suite
		router   *gin.Engine
		recorder *httptest.ResponseRecorder
	}
)

func (suite *apiTestSuite) SetupTest() {
	suite.router = getRouter()
	suite.recorder = httptest.NewRecorder()
}

func (suite *apiTestSuite) TestIndexAndPing() {
	// ping test
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	r := httptest.NewRecorder()
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(200, r.Code)
	suite.Assert().Equal("{}", r.Body.String())

	//index test
	req, _ = http.NewRequest(http.MethodGet, "/", nil)
	r = httptest.NewRecorder()
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(200, r.Code)
	suite.Assert().Equal("{}", r.Body.String())
}

func (suite *apiTestSuite) TestUser() {
	var user User
	newUser := []byte(`{"email":"test@email.com","name":"exampleUser","password":"dummy"}`)

	//create a new user
	req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(newUser))
	req.Header.Add("Content-Type", "application/json")
	r := httptest.NewRecorder()
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(200, r.Code)
	err := json.Unmarshal(r.Body.Bytes(), &user)
	suite.Assert().NoError(err)
	suite.Assert().NotEmpty(user)
	suite.Assert().Equal("test@email.com", user.Email)
	suite.Assert().Equal("exampleUser", user.Name)

	// get the user
	req, _ = http.NewRequest(http.MethodGet, "/user/"+user.Id, nil)
	r = httptest.NewRecorder()
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(200, r.Code)
	err = json.Unmarshal(r.Body.Bytes(), &user)
	suite.Assert().NoError(err)
	suite.Assert().Equal("test@email.com", user.Email)
	suite.Assert().Equal("exampleUser", user.Name)

	// get all users
	req, _ = http.NewRequest(http.MethodGet, "/users", nil)
	r = httptest.NewRecorder()
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(200, r.Code)
	suite.Assert().NotEmpty(r.Body.String())

	// update a user with new name and pass
	updateUser := []byte(`{"name":"exampleUser1","password":"dummy1"}`)
	req, _ = http.NewRequest(http.MethodPut, "/user/"+user.Id, bytes.NewBuffer(updateUser))
	r = httptest.NewRecorder()
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(200, r.Code)
	suite.Assert().NotEmpty(r.Body.String())

	// delete the user
	req, _ = http.NewRequest(http.MethodDelete, "/user/"+user.Id, nil)
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(200, r.Code)
}

func (suite *apiTestSuite) TestGroup() {
	var group Group
	var user User

	//create a new user
	newUser := []byte(`{"email":"test1@email.com","name":"exampleUser1","password":"dummy"}`)
	req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(newUser))
	req.Header.Add("Content-Type", "application/json")
	r := httptest.NewRecorder()
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(200, r.Code)
	err := json.Unmarshal(r.Body.Bytes(), &user)
	suite.Assert().NoError(err)

	// add a new group
	newGroup := []byte(`{"name":"testGroup"}`)
	req, _ = http.NewRequest(http.MethodPost, "/group", bytes.NewBuffer(newGroup))
	req.Header.Add("Content-Type", "application/json")
	r = httptest.NewRecorder()
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(201, r.Code)
	err = json.Unmarshal(r.Body.Bytes(), &group)
	suite.Assert().NoError(err)

	// get the group
	req, _ = http.NewRequest(http.MethodGet, "/group/"+group.Id, bytes.NewBuffer(newGroup))
	r = httptest.NewRecorder()
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(200, r.Code)

	// get all groups
	req, _ = http.NewRequest(http.MethodGet, "/groups", nil)
	r = httptest.NewRecorder()
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(200, r.Code)
	suite.Assert().NotEmpty(r.Body.String())

	// update group with a new name
	newName := []byte(`{"name":"testGroup1"}`)
	req, _ = http.NewRequest(http.MethodPut, "/group/"+group.Id, bytes.NewBuffer(newName))
	req.Header.Add("Content-Type", "application/json")
	r = httptest.NewRecorder()
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(200, r.Code)
	suite.Assert().NotEmpty(r.Body.String())

	// add a user to the group

	userToAdd := struct {
		UserId string `json:"userId"`
	}{UserId: user.Id}
	resBytes, err := json.Marshal(userToAdd)
	log.Println(userToAdd)
	suite.Assert().NoError(err)

	req, _ = http.NewRequest(http.MethodPost, "/group/"+group.Id+"/user", bytes.NewBuffer(resBytes))
	req.Header.Add("Content-Type", "application/json")
	r = httptest.NewRecorder()
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(200, r.Code)
	suite.Assert().NotEmpty(r.Body.String())

	// remove the user from the group
	req, _ = http.NewRequest(http.MethodDelete, "/group/"+group.Id+"/user/"+user.Id, nil)
	r = httptest.NewRecorder()
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(200, r.Code)

	// delete the group
	req, _ = http.NewRequest(http.MethodDelete, "/group/"+group.Id, nil)
	r = httptest.NewRecorder()
	suite.router.ServeHTTP(r, req)
	suite.Assert().Equal(200, r.Code)
}

func TestAPI(t *testing.T) {
	db.Connect(configuration.Mongo{
		Host:     "localhost",
		Username: "root",
		Password: "examplepass",
		Port:     27017,
		Database: "userManagement",
	})
	suite.Run(t, new(apiTestSuite))
}
