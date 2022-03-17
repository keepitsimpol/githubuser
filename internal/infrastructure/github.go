package infrastructure

import "github.com/keepitsimpol/githubuser/internal/core/gateway"

type githubClientImpl struct{}

func New() *githubClientImpl {
	return new(githubClientImpl)
}

func (c *githubClientImpl) GetGithubUser(user string) gateway.GetGithubUserResponse {
	return gateway.GetGithubUserResponse{}
}
