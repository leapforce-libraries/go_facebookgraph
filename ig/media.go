package facebookgraph

import (
	"fmt"
	"strings"

	fb2 "github.com/huandu/facebook/v2"
	errortools "github.com/leapforce-libraries/go_errortools"
	api "github.com/leapforce-libraries/go_facebookgraph/api"
)

type Media struct {
	ID            string `mapstructure:"id"`
	Timestamp     string `mapstructure:"timestamp"`
	MediaType     string `mapstructure:"media_type"`
	Permalink     string `mapstructure:"permalink"`
	Caption       string `mapstructure:"caption"`
	LikeCount     int64  `mapstructure:"like_count"`
	CommentsCount int64  `mapstructure:"comments_count"`
}

// Media returns Instagram media details
//
func (service *Service) Media(mediaID string, fields []string) (*Media, *errortools.Error) {
	path := fmt.Sprintf("/%s", mediaID)

	params := fb2.Params{
		"fields": strings.Join(fields, ","),
	}

	result, e := api.GetWithRetry(service.session, path, params)
	if e != nil {
		return nil, e
	}

	media := Media{}

	err := result.DecodeField("", &media)
	//err = mapstructure.Decode(result, &media)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	return &media, nil
}
