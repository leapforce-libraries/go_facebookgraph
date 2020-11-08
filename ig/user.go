package facebookgraph

import (
	"fmt"
	"strings"

	api "github.com/leapforce-libraries/go_facebookgraph/api"
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
func (ig *Instagram) User(userID string, fields []string) (*User, error) {
	path := fmt.Sprintf("/%s", userID)
	params := make(map[string]interface{})
	params["fields"] = strings.Join(fields, ",")

	result, err := api.GetWithRetry(ig.session, path, params)
	if err != nil {
		return nil, err
	}

	user := User{}

	err = result.DecodeField("", &user)
	//err = mapstructure.Decode(result, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
