package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/keepitsimpol/githubuser/internal/core/port"
	"github.com/keepitsimpol/githubuser/internal/core/util"
)

type githubClientImpl struct{}

func New() *githubClientImpl {
	return new(githubClientImpl)
}

type errorResponse struct {
	Message string `json:"message"`
}

const (
	BaseURL    = "https://api.github.com/users/"
	RetryCount = 3
	RetryWait  = 1 * time.Second
)

func (c *githubClientImpl) GetGithubUser(user string, ctx context.Context) (clientResponse port.GetGithubUserResponse, err error) {
	url := BaseURL + user
	util.Infof(ctx, "Start github client to get user detail with URL: %s", url)
	client := resty.New().
		SetRetryCount(RetryCount).
		SetRetryWaitTime(RetryWait)

	var errorResponse *errorResponse
	response, err := client.R().
		SetError(&errorResponse).
		SetResult(&clientResponse).
		Get(url)

	util.Infof(ctx, "Raw response: %+v", response)
	if err != nil {
		return clientResponse, err
	}

	if errorResponse != nil {
		return clientResponse, errors.New(fmt.Sprintf("error from github: %+v", errorResponse))
	}

	util.Infof(ctx, "End calling github with response: %+v", clientResponse)
	return clientResponse, err
}
