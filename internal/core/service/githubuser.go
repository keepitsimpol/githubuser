package service

import (
	"github.com/keepitsimpol/githubuser/internal/core/constant/errorcode"
	"github.com/keepitsimpol/githubuser/internal/core/model"
	"github.com/sirupsen/logrus"
)

type githubUserService struct {
	logger *logrus.Entry
}

func New(logger *logrus.Entry) *githubUserService {
	service := new(githubUserService)
	service.logger = logger
	return service
}

// GetGithubUsers Gets Github users' details and returns apprioriate app error code and error message
func (s githubUserService) GetGithubUsers(users []string) ([]model.GetGithubUserDetails, errorcode.AppErrorCode, error) {
	s.logger.Infof("Start get github users service with request: %v", users)
	return []model.GetGithubUserDetails{}, errorcode.NoError, nil
}
