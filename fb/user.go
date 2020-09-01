package facebookgraph

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type User struct {
	ID         string `mapstructure:"id"`
	Email      string `mapstructure:"email"`
	FirstName  string `mapstructure:"first_name"`
	HomeTown   string `mapstructure:"hometown"`
	LastName   string `mapstructure:"last_name"`
	MiddleName string `mapstructure:"middle_name"`
	Name       string `mapstructure:"name"`
	ProfilePic string `mapstructure:"profile_pic"`
	Website    string `mapstructure:"website"`
}

// User returns Facebook user details
//
func (fb *Facebook) User(userID string) (*User, error) {
	result, err := fb.session.Get(fmt.Sprintf("/%s", userID), nil)
	if err != nil {
		return nil, err
	}

	user := User{}

	err = mapstructure.Decode(result, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
