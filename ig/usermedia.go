package facebookgraph

import (
	"fmt"

	fb2 "github.com/huandu/facebook/v2"
	errortools "github.com/leapforce-libraries/go_errortools"
	api "github.com/leapforce-libraries/go_facebookgraph/api"
	models "github.com/leapforce-libraries/go_facebookgraph/models"
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
func (service *Service) UserMedia(userID string, after string) (*UserMediaResponse, *errortools.Error) {
	path := fmt.Sprintf("/%s/media", userID)

	params := fb2.Params{
		"limit": limit,
		"after": after,
	}

	result, e := api.GetWithRetry(service.session, path, params)
	if e != nil {
		return nil, e
	}

	response := UserMediaResponse{}

	err := result.DecodeField("", &response)
	//err = mapstructure.Decode(result, &response)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	return &response, nil
}
