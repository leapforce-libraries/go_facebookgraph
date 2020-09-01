package facebookgraph

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type User struct {
	ID                string `mapstructure:"id"`
	InstagramID       string `mapstructure:"ig_id"`
	Name              string `mapstructure:"name"`
	ProfilePictureURL string `mapstructure:"profile_picture_url"`
	UserName          string `mapstructure:"username"`
	Website           string `mapstructure:"website"`
	FollowersCount    int64  `mapstructure:"followers_count"`
	FollowsCount      int64  `mapstructure:"follows_count"`
	MediaCount        int64  `mapstructure:"media_count"`
}

// User returns Instagram user details
//
func (ig *FacebookGraph) User(userID string, fields []string) (*User, error) {
	params := make(map[string]interface{})
	params["fields"] = strings.Join(fields, ",")

	result, err := ig.session.Get(fmt.Sprintf("/%s", userID), params)
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
