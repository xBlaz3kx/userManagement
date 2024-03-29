package group

import "github.com/kamva/mgm/v3"

type (
	Group struct {
		mgm.DefaultModel `bson:",inline"`
		Name             string   `form:"name" binding:"required" json:"name" bson:"name" validate:"required"`
		Users            []string `json:"users" bson:"users"`
	}
)

//NewGroup creates a new group object
func NewGroup(name string) (*Group, error) {
	return &Group{
		Name:  name,
		Users: []string{},
	}, nil
}
