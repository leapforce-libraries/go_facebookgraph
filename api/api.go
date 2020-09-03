package facebookgraph

import (
	"fmt"
	"time"

	fb "github.com/huandu/facebook/v2"
	"github.com/mitchellh/mapstructure"

	models "github.com/Leapforce-nl/go_facebookgraph/models"
)

func getWithRetry(session *fb.Session, path string, params fb.Params) (fb.Result, error) {
	maxRetry := 10
	retry := 0
	var result fb.Result
	var err error

	for retry < maxRetry {
		result, err = session.Get(path, params)
		if err != nil {
			errorResponse := models.ErrorResponse{}
			err = mapstructure.Decode(result, &errorResponse)
			if err != nil {
				return nil, err
			}

			if errorResponse.Error.Code == 190 {
				retry++
				time.Sleep(3 * time.Second)
				fmt.Println("attempt:", retry)
			} else {
				return nil, err
			}
		}

		retry = maxRetry
	}

	return result, nil
}
