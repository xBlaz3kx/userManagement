package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/xBlaz3kx/userManagementExample/db"
	"github.com/xBlaz3kx/userManagementExample/internal/configuration"
	"github.com/xBlaz3kx/userManagementExample/internal/group"
	"github.com/xBlaz3kx/userManagementExample/internal/user"
	"net/http"
)

var pingHandler = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

func getRouter() *gin.Engine {
	router := gin.Default()

	// health check and index handler
	router.GET("/ping", pingHandler)
	router.GET("/", pingHandler)

	//users
	router.GET("/users", user.GetUsers)
	router.GET("/user/:id", user.GetUser)
	router.POST("/user", user.CreateNewUser)
	router.DELETE("/user/:id", user.DeleteUser)
	router.PUT("/user/:id", user.UpdateUser)
	router.PATCH("/user/:id", user.UpdateUser)

	//groups
	router.GET("/groups", group.GetGroups)
	router.GET("/group/:id", group.GetGroup)
	router.POST("/group", group.CreateNewGroup)
	router.DELETE("/group/:id", group.DeleteGroup)
	router.PUT("/group/:id", group.UpdateGroups)
	router.PATCH("/group/:id", group.UpdateGroups)

	router.POST("/group/:id/user", group.AddUserToGroup)
	router.DELETE("/group/:groupId/user/:userId", group.RemoveUserFromGroup)
	return router
}

func main() {
	router := getRouter()

	// get app configuration
	appConfiguration := configuration.GetConfiguration()

	api := appConfiguration.API
	connectAddress := fmt.Sprintf("%s:%d", api.Host, api.Port)

	// connect to the database
	db.Connect(appConfiguration.Mongo)

	// listen and serve the api on provided address
	err := endless.ListenAndServe(connectAddress, router)
	if err != nil {
		return
	}
}
