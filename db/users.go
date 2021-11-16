package db

import (
	"errors"
	"github.com/kamva/mgm/v3"
	. "github.com/xBlaz3kx/userManagementExample/model/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var ErrUserAlreadyExists = errors.New("user already exists")

func getUser(filter interface{}) (*User, error) {
	user := &User{}
	err := mgm.Coll(&User{}).First(filter, user)
	return user, err
}

func AddUser(user User) (*User, error) {
	log.Println("Adding a user", user.Name)
	// if the email is already used, skip
	_, err := getUser(bson.M{"email": user.Email})
	if err == mongo.ErrNoDocuments {
		// assume the password is in plain text and encrypt the password
		hash, encryptErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if encryptErr != nil {
			return nil, encryptErr
		}

		user.Password = string(hash)
		err = mgm.Coll(&User{}).Create(&user)
		return &user, err
	}

	return nil, ErrUserAlreadyExists
}

func GetUsers() ([]User, error) {
	log.Println("Getting all users..")
	var allUsers []User

	err := mgm.Coll(&User{}).SimpleFind(&allUsers, bson.M{})
	if err != nil {
		return nil, err
	}

	return allUsers, nil
}

func GetUser(id string) (*User, error) {
	log.Println("Getting a user", id)
	hex, _ := primitive.ObjectIDFromHex(id)
	return getUser(bson.M{"_id": hex})
}

func DeleteUser(id string) (bool, error) {
	log.Println("Deleting a user", id)
	hex, _ := primitive.ObjectIDFromHex(id)
	deleted, err := mgm.Coll(&User{}).DeleteOne(mgm.Ctx(), bson.M{"_id": hex})

	if deleted.DeletedCount > 0 {
		group, err := getGroup(bson.M{"users": hex})
		if err != nil {
			return false, err
		}

		RemoveUserFromGroup(group.ID.Hex(), id)
	}

	return deleted.DeletedCount > 0, err
}

func UpdateUser(id string, name, password *string) (*User, error) {
	log.Println("Updating a user", id)
	hex, _ := primitive.ObjectIDFromHex(id)

	user, err := getUser(bson.M{"_id": hex})
	if err != nil {
		return nil, err
	}

	if name != nil {
		user.Name = *name
	}

	if password != nil {
		// assume the password is in plain text and encrypt the password
		hash, encryptErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if encryptErr != nil {
			return nil, encryptErr
		}
		user.Password = string(hash)
	}

	err = mgm.Coll(&User{}).Update(user)

	if password != nil {
		user.Password = *password
	}

	return user, err
}
