package facebookgraph

import (
	"fmt"

	utils "github.com/Leapforce-nl/go_utilities"
	fb2 "github.com/huandu/facebook/v2"
)

type Page struct {
	ID          string   `mapstructure:"id"`
	AccessToken string   `mapstructure:"access_token"`
	Bio         string   `mapstructure:"bio"`
	Category    string   `mapstructure:"category"`
	Description string   `mapstructure:"description"`
	Emails      []string `mapstructure:"emails"`
	Name        string   `mapstructure:"name"`
	Website     string   `mapstructure:"website"`
}

// Page returns Facebook page details
//
func (fb *Facebook) Page(pageID string) (*Page, error) {

	path := fmt.Sprintf("/%s", pageID)
	params := fb2.Params{
		"fields": utils.GetTaggedTagNames("mapstructure", Page{}),
	}

	result, err := fb.session.Get(path, params)
	//result, err := fg.getWithRetry(path, params)

	page := Page{}

	err = result.DecodeField("", &page)
	//err = mapstructure.Decode(result, &page)
	if err != nil {
		return nil, err
	}

	return &page, nil
}
