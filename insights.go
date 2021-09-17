package facebookgraph

import (
	"fmt"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	api "github.com/leapforce-libraries/go_facebookgraph/api"
	models "github.com/leapforce-libraries/go_facebookgraph/models"
	"github.com/mitchellh/mapstructure"

	fb "github.com/huandu/facebook/v2"
)

type InsightsResponse struct {
	Data   []Insight     `mapstructure:"data"`
	Paging models.Paging `mapstructure:"paging"`
}

type InsightValue struct {
	Value   string `mapstructure:"value"`
	EndTime string `mapstructure:"end_time"`
}
type Insight struct {
	Name        string         `mapstructure:"name"`
	Period      string         `mapstructure:"period"`
	Values      []InsightValue `mapstructure:"values"`
	Title       string         `mapstructure:"title"`
	Description string         `mapstructure:"description"`
	ID          string         `mapstructure:"id"`
}

// Insights return Instagram insights for a user
//
func (service *Service) Insights(objectID string, metrics []string, period *string, since *int64, until *int64, accessToken *string) (*[]Insight, *fb.Error, *errortools.Error) {
	path := fmt.Sprintf("/%s/insights", objectID)

	params := fb.Params{}
	params["metric"] = strings.Join(metrics, ",")

	if period != nil {
		params["period"] = *period
	}
	if since != nil {
		params["since"] = *since
	}
	if until != nil {
		params["until"] = *until
	}
	if accessToken != nil {
		params["access_token"] = *accessToken
	}

	result, e := api.GetWithRetry(service.session, path, params)
	err := result.Err()
	if err != nil {
		if fbError, ok := err.(*fb.Error); ok {
			return nil, fbError, e
		}
	}
	if e != nil {
		return nil, nil, e
	}

	response := InsightsResponse{}
	err = mapstructure.Decode(result, &response)
	if err != nil {
		return nil, nil, errortools.ErrorMessage(err)
	}

	return &response.Data, nil, nil
}
