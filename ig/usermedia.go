package facebookgraph

import (
	"fmt"

	api "github.com/Leapforce-nl/go_facebookgraph/api"
	models "github.com/Leapforce-nl/go_facebookgraph/models"

	fb2 "github.com/huandu/facebook/v2"
)

const limit int = 100

type UserMediaResponse struct {
	Data   []UserMedia   `mapstructure:"data"`
	Paging models.Paging `mapstructure:"paging"`
}

type UserMedia struct {
	ID string `mapstructure:"id"`
}

// UserMedia return Instagram medias for a user
//
func (ig *Instagram) UserMedia(userID string, after string) (*UserMediaResponse, error) {
	path := fmt.Sprintf("/%s/media", userID)

	params := fb2.Params{
		"limit": limit,
		"after": after,
	}

	result, err := api.GetWithRetry(ig.session, path, params)
	if err != nil {
		return nil, err
	}

	response := UserMediaResponse{}

	err = result.DecodeField("", &response)
	//err = mapstructure.Decode(result, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
