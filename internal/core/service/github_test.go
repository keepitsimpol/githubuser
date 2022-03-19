package service

import (
	"context"
	"testing"

	"github.com/keepitsimpol/githubuser/internal/core/constant/errorcode"
	"github.com/keepitsimpol/githubuser/internal/core/mock"
	"github.com/keepitsimpol/githubuser/internal/core/model"
	"github.com/keepitsimpol/githubuser/internal/core/util"
	. "github.com/onsi/gomega"
)

func TestGetAccountDetails(t *testing.T) {
	g := NewGomegaWithT(t)
	scenarios := []struct {
		testcase          string
		usernames         []string
		expectedFirstName string
		shouldTestCache   bool
		hasError          bool
		hasClientError    bool
		errorType         errorcode.AppErrorCode
	}{
		{
			testcase:          "Request with 10 usernames",
			usernames:         []string{"nboy1", "bboy2", "aboy3", "aboy4", "eboy5", "fboy6", "gboy7", "hboy8", "iboy9", "jboy10"},
			expectedFirstName: "aboy3",
			errorType:         errorcode.NoError,
		},
		{
			testcase:          "Request with 10 usernames - test cache",
			usernames:         []string{"nboy1", "bboy2", "aboy3", "aboy4", "eboy5", "fboy6", "gboy7", "hboy8", "iboy9", "jboy10"},
			shouldTestCache:   true,
			expectedFirstName: "aboy3",
			errorType:         errorcode.NoError,
		},
		{
			testcase:  "Request with more than 10 usernames",
			usernames: []string{"nboy1", "bboy2", "aboy3", "aboy4", "eboy5", "fboy6", "gboy7", "hboy8", "iboy9", "jboy10", "zboy"},
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
					sampleResponse := response[0]
					g.Expect(sampleResponse.Name).To(Equal(tc.expectedFirstName))
				}
			}
			g.Expect(errorType).To(Equal(tc.errorType))
		})
	}
}
