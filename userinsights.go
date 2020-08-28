package instagramgraph

import (
	"fmt"
	"strings"

	fb "github.com/huandu/facebook/v2"
)

// UserInsights return Instagram insights for a user
//
func (ig *InstagramGraph) UserInsights(userID string, metrics []string, period string) (fb.Result, error) {
	params := make(map[string]interface{})
	//params["access_token"] = ig.accessToken
	params["metric"] = strings.Join(metrics, ",")
	params["period"] = period

	res2, err := ig.session.Get(fmt.Sprintf("/%s/insights", userID), params)
	if err != nil {
		return nil, err
	}

	return res2, nil
}
