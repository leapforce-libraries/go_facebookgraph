package facebookgraph

import (
	"fmt"

	"github.com/mitchellh/mapstructure"

	paging "github.com/Leapforce-nl/go_facebookgraph/paging"
)

const limit int = 100

type UserMediaResponse struct {
	Data   []UserMedia   `mapstructure:"data"`
	Paging paging.Paging `mapstructure:"paging"`
}

type UserMedia struct {
	ID string `mapstructure:"id"`
}

// UserMedia return Instagram medias for a user
//
func (ig *Instagram) UserMedia(userID string, after string) (*UserMediaResponse, error) {
	path := fmt.Sprintf("/%s/media?limit=%v&after=%s", userID, limit, after)

	result, err := ig.session.Get(path, nil)
	if err != nil {
		return nil, err
	}

	response := UserMediaResponse{}
	err = mapstructure.Decode(result, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
