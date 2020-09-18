package facebookgraph

import (
	"fmt"

	api "github.com/Leapforce-nl/go_facebookgraph/api"
	fb2 "github.com/huandu/facebook/v2"
)

const postCommentsLimit int = 50 //limit 100 icm comments does not work...

type PostCommentsResponse struct {
	Data    []PagePost          `mapstructure:"data"`
	Paging  Paging              `mapstructure:"paging"`
	Summary PostCommentsSummary `mapstructure:"paging"`
}

type PostCommentsSummary struct {
	Order      string `mapstructure:"order"`
	TotalCount int64  `mapstructure:"total_count"`
	CanComment bool   `mapstructure:"can_comment"`
}

type PostComment struct {
	ID          string `mapstructure:"id"`
	CreatedTime string `mapstructure:"created_time"`
	Message     string `mapstructure:"message"`
}

// PostComments returns Facebook post comments for a post
//
func (fb *Facebook) PostComments(postID string, accessToken string, after string) (*PostCommentsResponse, error) {
	path := fmt.Sprintf("/%s/comments", postID)

	params := fb2.Params{
		"limit":        postCommentsLimit,
		"after":        after,
		"access_token": accessToken,
		"summary":      false,
	}

	result, err := api.GetWithRetry(fb.session, path, params)
	if err != nil {
		return nil, err
	}

	response := PostCommentsResponse{}
	err = result.DecodeField("", &response)
	//err = mapstructure.Decode(result, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// PostCommentsCount returns Facebook post comments count for a post
//
func (fb *Facebook) PostCommentsCount(postID string, accessToken string) (*int64, error) {
	path := fmt.Sprintf("/%s/comments", postID)

	params := fb2.Params{
		"limit":        0,
		"access_token": accessToken,
		"summary":      true,
	}

	result, err := api.GetWithRetry(fb.session, path, params)
	if err != nil {
		return nil, err
	}

	response := PostCommentsResponse{}
	err = result.DecodeField("", &response)
	//err = mapstructure.Decode(result, &response)
	if err != nil {
		return nil, err
	}

	return &response.Summary.TotalCount, nil
}
