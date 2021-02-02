package facebookgraph

import (
	"fmt"
	"time"

	fb "github.com/huandu/facebook/v2"
	"github.com/mitchellh/mapstructure"

	errortools "github.com/leapforce-libraries/go_errortools"
	models "github.com/leapforce-libraries/go_facebookgraph/models"
)

const (
	maxRetry          int   = 10
	errorCodeRetry    int64 = 190
	retryWaitXSeconds int   = 3
)

func GetWithRetry(session *fb.Session, path string, params fb.Params) (fb.Result, *errortools.Error) {
	retry := 0
	var result fb.Result
	var err error

	for retry < maxRetry {
		result, err = session.Get(path, params)
		if err != nil {
			errorResponse := models.ErrorResponse{}
			err2 := mapstructure.Decode(result, &errorResponse)
			if err2 != nil {
				return nil, errortools.ErrorMessage(err2)
			}

			if errorResponse.Error.Code == errorCodeRetry {
				retry++
				time.Sleep(time.Duration(retryWaitXSeconds) * time.Second)
				fmt.Println("attempt:", retry)
			} else {
				return nil, errortools.ErrorMessage(err)
			}
		}

		retry = maxRetry
	}

	return result, nil
}
