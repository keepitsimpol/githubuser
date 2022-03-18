package service

import (
	"context"
	"testing"

	"github.com/keepitsimpol/githubuser/internal/core/constant/errorcode"
	"github.com/keepitsimpol/githubuser/internal/core/mock"
	"github.com/keepitsimpol/githubuser/internal/core/model"
	"github.com/keepitsimpol/githubuser/internal/core/port"
	"github.com/keepitsimpol/githubuser/internal/core/util"
	. "github.com/onsi/gomega"
)

func TestGetAccountDetails(t *testing.T) {
	g := NewGomegaWithT(t)
	scenarios := []struct {
		testcase        string
		usernames       []string
		response        port.GetGithubUserResponse
		shouldTestCache bool
		hasError        bool
		hasClientError  bool
		errorType       errorcode.AppErrorCode
	}{
		{
			testcase:  "Request with 10 usernames",
			usernames: []string{"boy1", "boy2", "boy3", "boy4", "boy5", "boy6", "boy7", "boy8", "boy9", "boy10"},
			response: port.GetGithubUserResponse{
				Name:        "Boy 1",
				Login:       "Boy 1",
				Company:     "ABC",
				Followers:   0,
				PublicRepos: 0,
			},
			errorType: errorcode.NoError,
		},
		{
			testcase:        "Request with 10 usernames - test cache",
			usernames:       []string{"boy1", "boy2", "boy3", "boy4", "boy5", "boy6", "boy7", "boy8", "boy9", "boy10"},
			shouldTestCache: true,
			response: port.GetGithubUserResponse{
				Name:        "Boy 1",
				Login:       "Boy 1",
				Company:     "ABC",
				Followers:   0,
				PublicRepos: 0,
			},
			errorType: errorcode.NoError,
		},
		{
			testcase:  "Request with more than 10 usernames",
			usernames: []string{"boy1", "boy2", "boy3", "boy4", "boy5", "boy6", "boy7", "boy8", "boy9", "boy10", "boy11"},
			hasError:  true,
			errorType: errorcode.InvalidRequest,
		},
		{
			testcase:  "Request is empty",
			usernames: []string{},
			hasError:  true,
			errorType: errorcode.InvalidRequest,
		},
		{
			testcase:       "Has Github client error",
			usernames:      []string{"boy1"},
			hasError:       false,
			hasClientError: true,
			errorType:      errorcode.NoError,
		},
	}

	for _, tc := range scenarios {
		t.Run(tc.testcase, func(t *testing.T) {
			util.GetCache().Clear()
			mockGithubClient := mock.Builder().
				MockGetGithubUserResponse(tc.response).
				GetGithubUserHasError(tc.hasClientError).
				Build()

			service := New(mockGithubClient)
			response, errorType, err := service.GetAccountDetails(model.GetAccountDetailRequest{UserNames: tc.usernames},
				context.Background())

			if tc.shouldTestCache {
				response, errorType, err = service.GetAccountDetails(model.GetAccountDetailRequest{UserNames: tc.usernames},
					context.Background())
			}

			if tc.hasError {
				g.Expect(err).ToNot(BeNil())
			} else {
				g.Expect(err).To(BeNil())

				if !tc.hasClientError {
					sampleResponse := response[0].(port.GetGithubUserResponse)
					g.Expect(sampleResponse.Name).To(Equal(tc.response.Name))
					g.Expect(sampleResponse.Login).To(Equal(tc.response.Login))
					g.Expect(sampleResponse.Company).To(Equal(tc.response.Company))
					g.Expect(sampleResponse.Followers).To(Equal(tc.response.Followers))
					g.Expect(sampleResponse.PublicRepos).To(Equal(tc.response.PublicRepos))
				}
			}
			g.Expect(errorType).To(Equal(tc.errorType))
		})
	}
}
