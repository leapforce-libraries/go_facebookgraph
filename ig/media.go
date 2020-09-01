package facebookgraph

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type Media struct {
	ID        string `mapstructure:"id"`
	Timestamp string `mapstructure:"timestamp"`
	MediaType string `mapstructure:"media_type"`
	MediaURL  string `mapstructure:"media_url"`
	Caption   string `mapstructure:"caption"`
	LikeCount int64  `mapstructure:"like_count"`
}

// Media returns Instagram media details
//
func (ig *FacebookGraph) Media(mediaID string, fields []string) (*Media, error) {
	params := make(map[string]interface{})
	params["fields"] = strings.Join(fields, ",")

	result, err := ig.session.Get(fmt.Sprintf("/%s", mediaID), params)
	if err != nil {
		return nil, err
	}

	media := Media{}

	err = mapstructure.Decode(result, &media)
	if err != nil {
		return nil, err
	}

	return &media, nil
}
