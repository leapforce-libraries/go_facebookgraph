package facebookgraph

import (
	"fmt"
	"strings"

	api "github.com/leapforce-libraries/go_facebookgraph/api"

	fb2 "github.com/huandu/facebook/v2"
)

type Media struct {
	ID            string `mapstructure:"id"`
	Timestamp     string `mapstructure:"timestamp"`
	MediaType     string `mapstructure:"media_type"`
	MediaURL      string `mapstructure:"media_url"`
	Caption       string `mapstructure:"caption"`
	LikeCount     int64  `mapstructure:"like_count"`
	CommentsCount int64  `mapstructure:"comments_count"`
}

// Media returns Instagram media details
//
func (ig *Instagram) Media(mediaID string, fields []string) (*Media, error) {
	path := fmt.Sprintf("/%s", mediaID)

	params := fb2.Params{
		"fields": strings.Join(fields, ","),
	}

	result, err := api.GetWithRetry(ig.session, path, params)
	if err != nil {
		return nil, err
	}

	media := Media{}

	err = result.DecodeField("", &media)
	//err = mapstructure.Decode(result, &media)
	if err != nil {
		return nil, err
	}

	return &media, nil
}
