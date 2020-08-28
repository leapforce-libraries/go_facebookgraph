package instagramgraph

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type Response struct {
	Data   []UserInsight `mapstructure:"data"`
	Paging Paging        `mapstructure:"paging"`
}

type Paging struct {
	Previous string `mapstructure:"previous"`
	Next     string `mapstructure:"next"`
}

type InsightValue struct {
	Value   string `mapstructure:"value"`
	EndTime string `mapstructure:"end_time"`
}
type UserInsight struct {
	Name        string         `mapstructure:"name"`
	Period      string         `mapstructure:"period"`
	Values      []InsightValue `mapstructure:"values"`
	Title       string         `mapstructure:"title"`
	Description string         `mapstructure:"description"`
	ID          string         `mapstructure:"id"`
}

// UserInsights return Instagram insights for a user
//
func (ig *InstagramGraph) UserInsights(userID string, metrics []string, period string, since int64, until int64) (insights *[]UserInsight, err error) {
	params := make(map[string]interface{})
	params["metric"] = strings.Join(metrics, ",")
	params["period"] = period
	params["since"] = since
	params["until"] = until

	result, err := ig.session.Get(fmt.Sprintf("/%s/insights", userID), params)
	if err != nil {
		return nil, err
	}

	response := Response{}
	err = mapstructure.Decode(result, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}
