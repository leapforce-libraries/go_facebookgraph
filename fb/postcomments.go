package facebookgraph

import (
	"fmt"

	fb2 "github.com/huandu/facebook/v2"
	errortools "github.com/leapforce-libraries/go_errortools"
	api "github.com/leapforce-libraries/go_facebookgraph/api"
)

type PostCommentsFilter string

const (
	PostCommentsFilterTopLevel PostCommentsFilter = "toplevel"
	PostCommentsFilterStream   PostCommentsFilter = "stream"
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
func (service *Service) PostComments(postID string, accessToken string, after string, filter *PostCommentsFilter) (*PostCommentsResponse, *errortools.Error) {
	path := fmt.Sprintf("/%s/comments", postID)

	params := fb2.Params{
		"limit":        postCommentsLimit,
		"after":        after,
		"access_token": accessToken,
		"summary":      false,
	}

	if filter != nil {
		params["filter"] = *filter
	}

	result, e := api.GetWithRetry(service.session, path, params)
	if e != nil {
		return nil, e
	}

	response := PostCommentsResponse{}
	err := result.DecodeField("", &response)
	//err = mapstructure.Decode(result, &response)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	return &response, nil
}

// PostCommentsCount returns Facebook post comments count for a post
//
func (service *Service) PostCommentsCount(postID string, accessToken string, filter *PostCommentsFilter) (*int64, *errortools.Error) {
	path := fmt.Sprintf("/%s/comments", postID)

	params := fb2.Params{
		"limit":        0,
		"access_token": accessToken,
		"summary":      true,
	}

	if filter != nil {
		params["filter"] = *filter
	}

	result, e := api.GetWithRetry(service.session, path, params)
	if e != nil {
		return nil, e
	}

	response := PostCommentsResponse{}
	err := result.DecodeField("", &response)
	//err = mapstructure.Decode(result, &response)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	return &response.Summary.TotalCount, nil
}
