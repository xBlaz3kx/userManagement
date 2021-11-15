package db

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	. "github.com/xBlaz3kx/userManagementExample/model/group"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var ErrUserInGroup = fmt.Errorf("user already in a group")
var ErrUserNotInGroup = fmt.Errorf("user not in the group")

func getGroup(filter interface{}) (*Group, error) {
	group := &Group{}
	err := mgm.Coll(&Group{}).First(filter, group)
	return group, err
}

func AddGroup(group Group) (*Group, error) {
	log.Println("Adding a group", group.Name)
	// check if group with the same name exists
	_, err := getGroup(bson.M{"name": group.Name})
	if err != nil && err == mongo.ErrNoDocuments {
		err = mgm.Coll(&Group{}).Create(&group)
		return &group, err
	}

	return &group, err
}

func GetGroups() ([]Group, error) {
	log.Println("Getting all groups..")
	var allGroups = []Group{}

	err := mgm.Coll(&Group{}).SimpleFind(&allGroups, bson.M{})
	if err != nil {
		return nil, err
	}

	return allGroups, nil
}

func GetGroup(id string) (*Group, error) {
	log.Println("Getting a group", id)
	hex, _ := primitive.ObjectIDFromHex(id)
	return getGroup(bson.M{"_id": hex})
}

func DeleteGroup(id string) (bool, error) {
	log.Println("Deleting a group", id)
	hex, _ := primitive.ObjectIDFromHex(id)
	result, err := mgm.Coll(&Group{}).DeleteOne(mgm.Ctx(), bson.M{"_id": hex})
	return result.DeletedCount > 0, err
}

func UpdateGroup(id string, newName *string) (*Group, error) {
	log.Println("Updating a group", id)
	hex, _ := primitive.ObjectIDFromHex(id)

	group, err := getGroup(bson.M{"_id": hex})
	if err != nil {
		return nil, err
	}

	if newName != nil {
		group.Name = *newName
	}
	err = mgm.Coll(&Group{}).Update(group)

	return group, err
}

func IsUserNotInGroup(groupId, userId string) bool {
	log.Printf("Checking if user %s is in a group %s", userId, groupId)
	hex, _ := primitive.ObjectIDFromHex(groupId)

	_, err := getGroup(bson.M{
		"_id":   hex,
		"users": userId,
	})
	log.Println(err)
	return err == mongo.ErrNoDocuments
}

func AddUserToGroup(groupId, userId string) (*Group, error) {
	log.Printf("Adding the user %s to the group %s", userId, groupId)
	if IsUserNotInGroup(groupId, userId) {
		hex, _ := primitive.ObjectIDFromHex(groupId)
		group, err := getGroup(bson.M{"_id": hex})
		if err != nil {
			return nil, err
		}

		group.Users = append(group.Users, userId)
		err = mgm.Coll(&Group{}).Update(group)
		return group, err
	}
	return nil, ErrUserInGroup
}

func RemoveUserFromGroup(groupId, userId string) (*Group, error) {
	log.Printf("Removing the user %s from the group %s", userId, groupId)
	if !IsUserNotInGroup(groupId, userId) {
		// get group with id
		hex, _ := primitive.ObjectIDFromHex(groupId)
		group, err := getGroup(bson.M{"_id": hex})
		if err != nil {
			return nil, err
		}

		for i, user := range group.Users {
			if user == userId {
				if i-1 < len(group.Users) {
					group.Users = append(group.Users[:i], group.Users[i+1:]...)
				} else {
					group.Users = append(group.Users[:i])
				}

			}
		}

		err = mgm.Coll(&Group{}).Update(group)
		return group, err

	}
	return nil, ErrUserNotInGroup
}
