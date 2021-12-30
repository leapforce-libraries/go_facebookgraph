package facebookgraph

import (
	"fmt"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	api "github.com/leapforce-libraries/go_facebookgraph/api"
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
func (service *Service) User(userID string, fields []string) (*User, *errortools.Error) {
	path := fmt.Sprintf("/%s", userID)
	params := make(map[string]interface{})
	params["fields"] = strings.Join(fields, ",")

	result, e := api.GetWithRetry(service.session, path, params)
	if e != nil {
		return nil, e
	}

	user := User{}

	err := mapstructure.Decode(result, &user)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	return &user, nil
}
