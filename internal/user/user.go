package user

import (
	"github.com/gin-gonic/gin"
	"github.com/xBlaz3kx/userManagementExample/db"
	"github.com/xBlaz3kx/userManagementExample/model/user"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type UpdateUserBody struct {
	Name     *string `binding:"required" json:"name"`
	Password *string `binding:"required" json:"password"`
}

type CreateUserBody struct {
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
	Name     string `binding:"required" json:"name"`
}

var CreateNewUser = func(context *gin.Context) {
	var body CreateUserBody

	if err := context.Bind(&body); err != nil {
		log.Println(body)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, _ := user.NewUser(body.Email, body.Password, body.Password)

	user, err := db.AddUser(*u)
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

	context.JSON(http.StatusOK, gin.H{"id": user.ID.Hex(), "email": user.Email, "name": user.Name, "password": user.Password})
}

var GetUser = func(context *gin.Context) {
	id := context.Param("id")

	user, err := db.GetUser(id)
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

	context.JSON(http.StatusOK, gin.H{"id": user.ID.Hex(), "email": user.Email, "name": user.Name, "password": user.Password})
}

var GetUsers = func(context *gin.Context) {
	users, err := db.GetUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": users})
}

var UpdateUser = func(context *gin.Context) {
	var updateBody UpdateUserBody

	id := context.Param("id")

	if err := context.ShouldBindJSON(&updateBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := db.UpdateUser(id, updateBody.Name, updateBody.Password)
	switch err {
	case nil:
		context.JSON(http.StatusOK, gin.H{"id": user.ID.Hex(), "name": user.Name, "password": user.Password, "email": user.Email})
		return
	case mongo.ErrNoDocuments:
		context.JSON(http.StatusNotFound, gin.H{})
		return
	default:
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

var DeleteUser = func(context *gin.Context) {
	id := context.Param("id")
	isDeleted, err := db.DeleteUser(id)

	switch err {
	case nil:
		if isDeleted {
			context.JSON(http.StatusOK, gin.H{})
			return
		}
		context.JSON(http.StatusNotFound, gin.H{})
		return
	case mongo.ErrNoDocuments:
		context.JSON(http.StatusNotFound, gin.H{})
		return
	default:
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
