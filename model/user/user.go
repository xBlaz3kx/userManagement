package user

import "github.com/kamva/mgm/v3"

type (
	User struct {
		mgm.DefaultModel `bson:",inline"`
		Email            string `json:"email" bson:"email" validate:"required,email"`
		Password         string `json:"password" bson:"password" validate:"required"`
		Name             string `json:"name"  bson:"name" validate:"required"`
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
