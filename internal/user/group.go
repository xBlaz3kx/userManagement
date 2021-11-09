package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var CreateNewUser = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

var GetUser = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

var GetUsers = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

var UpdateUser = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

var DeleteUser = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}
