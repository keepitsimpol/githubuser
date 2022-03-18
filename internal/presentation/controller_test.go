package presentation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/keepitsimpol/githubuser/internal/core/mock"
	"github.com/keepitsimpol/githubuser/internal/core/port"
	"github.com/keepitsimpol/githubuser/internal/core/service"
	"github.com/keepitsimpol/githubuser/internal/core/util"
	. "github.com/onsi/gomega"
)

func TestGetUserAccountDetails(t *testing.T) {
	g := NewGomegaWithT(t)
	scenarios := []struct {
		testcase        string
		usernames       []string
		response        port.GetGithubUserResponse
		result          bool
		message         string
		httpCode        int
		shouldTestCache bool
		hasClientError  bool
	}{
		{
			testcase:  "Request with 10 usernames",
			usernames: []string{"boy1", "boy2", "boy3", "boy4", "boy5", "boy6", "boy7", "boy8", "boy9", "boy10"},
			result:    true,
			httpCode:  http.StatusOK,
			response: port.GetGithubUserResponse{
				Name:        "Boy 1",
				Login:       "Boy 1",
				Company:     "ABC",
				Followers:   0,
				PublicRepos: 0,
			},
		},
		// {
		// 	testcase:        "Request with 10 usernames - test cache",
		// 	usernames:       []string{"boy1", "boy2", "boy3", "boy4", "boy5", "boy6", "boy7", "boy8", "boy9", "boy10"},
		// 	shouldTestCache: true,
		// 	response: port.GetGithubUserResponse{
		// 		Name:        "Boy 1",
		// 		Login:       "Boy 1",
		// 		Company:     "ABC",
		// 		Followers:   0,
		// 		PublicRepos: 0,
		// 	},
		// 	errorType: errorcode.NoError,
		// },
		// {
		// 	testcase:  "Request with more than 10 usernames",
		// 	usernames: []string{"boy1", "boy2", "boy3", "boy4", "boy5", "boy6", "boy7", "boy8", "boy9", "boy10", "boy11"},
		// 	hasError:  true,
		// 	errorType: errorcode.InvalidRequest,
		// },
		// {
		// 	testcase:  "Request is empty",
		// 	usernames: []string{},
		// 	hasError:  true,
		// 	errorType: errorcode.InvalidRequest,
		// },
		// {
		// 	testcase:       "Has Github client error",
		// 	usernames:      []string{"boy1"},
		// 	hasError:       false,
		// 	hasClientError: true,
		// 	errorType:      errorcode.NoError,
		// },
	}

	for _, tc := range scenarios {
		t.Run(tc.testcase, func(t *testing.T) {
			util.GetCache().Clear()
			mockGithubClient := mock.Builder().
				MockGetGithubUserResponse(tc.response).
				GetGithubUserHasError(tc.hasClientError).
				Build()

			serviceImpl := service.New(mockGithubClient)
			controller := New(serviceImpl)

			router := gin.Default()
			router.POST("/test", controller.GetUserAccountDetails)

			request := GetAccountDetailsRequest{Users: tc.usernames}
			reqBytes, err := json.Marshal(request)
			req, err := http.NewRequest("POST", "/test", bytes.NewReader(reqBytes))
			g.Expect(err).To(BeNil())

			writer := httptest.NewRecorder()
			router.ServeHTTP(writer, req)

			var response GetAccountDetailsResponse
			err = json.NewDecoder(writer.Body).Decode(&response)
			g.Expect(err).To(BeNil())

			if tc.shouldTestCache {
				req, err := http.NewRequest("POST", "/test", bytes.NewReader(reqBytes))
				g.Expect(err).To(BeNil())

				writer := httptest.NewRecorder()
				router.ServeHTTP(writer, req)
			}

			g.Expect(response.Result).To(Equal(tc.result))
			g.Expect(response.Message).To(Equal(tc.message))
			g.Expect(writer.Code).To(Equal(tc.httpCode))

			if tc.result {
				sampleResponse := response.UserDetails[0]
				g.Expect(sampleResponse.Name).To(Equal(tc.response.Name))
				g.Expect(sampleResponse.Login).To(Equal(tc.response.Login))
				g.Expect(sampleResponse.Company).To(Equal(tc.response.Company))
				g.Expect(sampleResponse.Followers).To(Equal(tc.response.Followers))
				g.Expect(sampleResponse.PublicRepos).To(Equal(tc.response.PublicRepos))
			}
		})
	}
}
