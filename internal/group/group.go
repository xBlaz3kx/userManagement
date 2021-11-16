package group

import (
	"github.com/gin-gonic/gin"
	"github.com/xBlaz3kx/userManagementExample/db"
	group2 "github.com/xBlaz3kx/userManagementExample/model/group"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type UserToGroup struct {
	UserId string `json:"userId"`
}

var CreateNewGroup = func(context *gin.Context) {
	var newGroup group2.Group

	if err := context.ShouldBindJSON(&newGroup); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group, err := db.AddGroup(newGroup)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"id": group.ID.Hex(), "name": newGroup.Name})
}

var GetGroups = func(context *gin.Context) {
	groups, err := db.GetGroups()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"group": groups})
}

var GetGroup = func(context *gin.Context) {
	id := context.Param("id")

	group, err := db.GetGroup(id)
	switch err {
	case nil:
		break
	case mongo.ErrNoDocuments:
		context.JSON(http.StatusNotFound, gin.H{})
		return
	default:
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"id": group.ID.Hex(), "name": group.Name})
}

var UpdateGroups = func(context *gin.Context) {
	var newGroup group2.Group
	id := context.Param("id")

	if err := context.ShouldBindJSON(&newGroup); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group, err := db.UpdateGroup(id, &newGroup.Name)
	switch err {
	case nil:
		break
	case mongo.ErrNoDocuments:
		context.JSON(http.StatusNotFound, gin.H{})
		return
	default:
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"id": group.ID.Hex(), "name": group.Name})
}

var DeleteGroup = func(context *gin.Context) {
	id := context.Param("id")

	isDeleted, err := db.DeleteGroup(id)
	switch err {
	case nil:
		if isDeleted {
			context.JSON(http.StatusOK, gin.H{})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	case mongo.ErrNoDocuments:
		context.JSON(http.StatusNotFound, gin.H{})
		return
	default:
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

var AddUserToGroup = func(context *gin.Context) {
	var user UserToGroup
	id := context.Param("id")

	if err := context.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group, err := db.AddUserToGroup(id, user.UserId)
	switch err {
	case nil:
		context.JSON(http.StatusOK, gin.H{"id": group.ID.Hex(), "name": group.Name, "users": group.Users})
		return
	case mongo.ErrNoDocuments:
		context.JSON(http.StatusNotFound, gin.H{})
		return
	case db.ErrUserInGroup:
		context.JSON(http.StatusNotModified, gin.H{})
		return
	default:
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

var RemoveUserFromGroup = func(context *gin.Context) {
	groupId := context.Param("id")
	userId := context.Param("userId")

	_, err := db.RemoveUserFromGroup(groupId, userId)
	switch err {
	case nil:
		context.JSON(http.StatusOK, gin.H{})
		return
	case mongo.ErrNoDocuments:
		context.JSON(http.StatusNotFound, gin.H{})
		return
	case db.ErrUserNotInGroup:
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	default:
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
