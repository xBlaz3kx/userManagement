package group

type (
	Group struct {
		Name  string   `json:"name" bson:"name" validate:"required"`
		Users []string `json:"users" bson:"users"`
	}
)

//NewGroup creates a new group object
func NewGroup(name string) (*Group, error) {
	return &Group{
		Name:  name,
		Users: []string{},
	}, nil
}
