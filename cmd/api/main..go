package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xBlaz3kx/userManagementExample/internal/group"
	"github.com/xBlaz3kx/userManagementExample/internal/user"
	"log"
	"net/http"
)

var pingHandler = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

func getRouter() *gin.Engine {
	router := gin.Default()

	// health check
	router.GET("/ping", pingHandler)

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
	router.DELETE("/group/:id/user", group.RemoveUserFromGroup)
	return router
}

func main() {
	router := getRouter()

	err := router.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Println(err)
	}
}
