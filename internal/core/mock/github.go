package mock

import (
	"context"
	"errors"

	"github.com/keepitsimpol/githubuser/internal/core/port"
)

type githubClientMock struct {
	GetGithubUserResponse port.GetGithubUserResponse
	hasError              bool
}

func (c *githubClientMock) GetGithubUser(_ string, _ context.Context) (clientResponse port.GetGithubUserResponse, err error) {
	if c.hasError {
		return port.GetGithubUserResponse{}, errors.New("mock GetGithubUser error")
	}
	return c.GetGithubUserResponse, nil
}

// Builder
type githubMockBuilder struct {
	clientMock *githubClientMock
}

func Builder() *githubMockBuilder {
	clientMock := new(githubClientMock)
	builder := new(githubMockBuilder)
	builder.clientMock = clientMock
	return builder
}

func (b *githubMockBuilder) MockGetGithubUserResponse(mockResponse port.GetGithubUserResponse) *githubMockBuilder {
	b.clientMock.GetGithubUserResponse = mockResponse
	return b
}

func (b *githubMockBuilder) GetGithubUserHasError(hasError bool) *githubMockBuilder {
	b.clientMock.hasError = hasError
	return b
}

func (b *githubMockBuilder) Build() *githubClientMock {
	return b.clientMock
}
