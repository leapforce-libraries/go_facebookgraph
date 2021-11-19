package facebookgraph

import (
	"fmt"

	fb2 "github.com/huandu/facebook/v2"
	errortools "github.com/leapforce-libraries/go_errortools"
	api "github.com/leapforce-libraries/go_facebookgraph/api"
	utils "github.com/leapforce-libraries/go_utilities"
)

type Page struct {
	ID             string   `mapstructure:"id"`
	AccessToken    string   `mapstructure:"access_token"`
	Bio            string   `mapstructure:"bio"`
	Category       string   `mapstructure:"category"`
	Description    string   `mapstructure:"description"`
	Emails         []string `mapstructure:"emails"`
	FanCount       uint32   `mapstructure:"fan_count"`
	FollowersCount uint32   `mapstructure:"followers_count"`
	Name           string   `mapstructure:"name"`
	Website        string   `mapstructure:"website"`
}

// Page returns Facebook page details
//
func (service *Service) Page(pageID string) (*Page, *errortools.Error) {

	path := fmt.Sprintf("/%s", pageID)
	params := fb2.Params{
		"fields": utils.GetTaggedTagNames("mapstructure", Page{}),
	}

	result, e := api.GetWithRetry(service.session, path, params)
	if e != nil {
		return nil, e
	}

	page := Page{}

	err := result.DecodeField("", &page)
	//err = mapstructure.Decode(result, &page)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	return &page, nil
}
