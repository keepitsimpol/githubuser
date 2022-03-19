package service

import (
	"errors"
	"fmt"

	"github.com/keepitsimpol/githubuser/internal/core/constant"
	"github.com/keepitsimpol/githubuser/internal/core/port"
)

type accountServiceFactoryImpl struct {
	githubService port.AccountDetailService
}

func NewAccountDetailServiceFactory(githubClient port.GithubClient) *accountServiceFactoryImpl {
	factory := new(accountServiceFactoryImpl)
	factory.githubService = NewAccountService(githubClient)
	return factory
}

func (a accountServiceFactoryImpl) GetAccountDetailService(source string) (port.AccountDetailService, error) {
	if source == constant.SourceGithub {
		return a.githubService, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Unsupported source: %s", source))
	}
}
