package user

import "github.com/kamva/mgm/v3"

type (
	User struct {
		mgm.DefaultModel `bson:",inline"`
		Email            string `form:"email" binding:"required" json:"email" bson:"email" validate:"required,email"`
		Password         string `form:"password" binding:"required" json:"password" bson:"password" validate:"required"`
		Name             string `form:"name" binding:"required" json:"name"  bson:"name" validate:"required"`
	}
)

//NewUser creates a new user object
func NewUser(email, name, password string) (*User, error) {
	return &User{
		Email:    email,
		Password: name,
		Name:     password,
	}, nil
}
