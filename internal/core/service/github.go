package service

import (
	"context"
	"errors"

	validate "github.com/go-playground/validator/v10"
	"github.com/keepitsimpol/githubuser/internal/core/constant/errorcode"
	"github.com/keepitsimpol/githubuser/internal/core/model"
	"github.com/keepitsimpol/githubuser/internal/core/port"
	"github.com/keepitsimpol/githubuser/internal/core/util"
)

type githubAccountDetailService struct {
	githubClient port.GithubClient
}

func New(githubClient port.GithubClient) *githubAccountDetailService {
	service := new(githubAccountDetailService)
	service.githubClient = githubClient
	return service
}

// GetAccountDetails Gets Github user details and returns apprioriate app error code and error message
func (s githubAccountDetailService) GetAccountDetails(users model.GetAccountDetailRequest, ctx context.Context) (response []model.GetAccountDetailResponse, appError errorcode.AppErrorCode, err error) {
	util.Infof(ctx, "Start getting github users service with request: %v", users)

	err = validate.New().Struct(users)
	if err != nil {
		util.Errorf(ctx, "Validation error in request: %s", err.Error())
		return response, errorcode.InvalidRequest, errors.New("request is invalid")
	}

	for _, user := range users.UserNames {
		if cachedUser, ok := util.GetCache().Get(user); ok {
			util.Infof(ctx, "User: %s is already cached. Reusing cached entry.", user)
			response = append(response, cachedUser.(model.GetAccountDetailResponse))
			continue
		}

		githubClientResponse, err := s.githubClient.GetGithubUser(user, ctx)
		if err != nil {
			util.Errorf(ctx, "Error while calling github for user: %s error: %s", user, err.Error())
			continue
		}

		util.Infof(ctx, "Adding account details for user: %s", user)
		response = append(response, model.GetAccountDetailResponse{
			Name:        githubClientResponse.Name,
			Login:       githubClientResponse.Login,
			Company:     githubClientResponse.Company,
			Followers:   githubClientResponse.Followers,
			PublicRepos: githubClientResponse.PublicRepos,
		})
		util.AddtoCache(user, githubClientResponse)
	}

	util.Infof(ctx, "End getting github users service with response: %+v", response)
	return response, errorcode.NoError, nil
}
