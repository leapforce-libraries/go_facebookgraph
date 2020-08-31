package instagramgraph

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type InsightsResponse struct {
	Data   []Insight `mapstructure:"data"`
	Paging Paging    `mapstructure:"paging"`
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
func (ig *InstagramGraph) Insights(objectID string, metrics []string, period string, since int64, until int64) (*[]Insight, error) {
	params := make(map[string]interface{})
	params["metric"] = strings.Join(metrics, ",")
	params["period"] = period
	params["since"] = since
	params["until"] = until

	result, err := ig.session.Get(fmt.Sprintf("/%s/insights", objectID), params)
	if err != nil {
		return nil, err
	}

	response := InsightsResponse{}
	err = mapstructure.Decode(result, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}
