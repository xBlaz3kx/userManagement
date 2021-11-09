package group

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var CreateNewGroup = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

var GetGroups = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

var GetGroup = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

var UpdateGroups = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

var DeleteGroup = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

var AddUserToGroup = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

var RemoveUserFromGroup = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}
