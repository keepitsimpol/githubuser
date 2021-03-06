package presentation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/keepitsimpol/githubuser/internal/core/mock"
	"github.com/keepitsimpol/githubuser/internal/core/service"
	"github.com/keepitsimpol/githubuser/internal/core/util"
	. "github.com/onsi/gomega"
)

func TestGetUserAccountDetails(t *testing.T) {
	g := NewGomegaWithT(t)
	scenarios := []struct {
		testcase                  string
		usernames                 []string
		expectedFirstName         string
		url                       string
		result                    bool
		message                   string
		httpCode                  int
		shouldTestCache           bool
		hasClientError            bool
		shouldErrorParsingRequest bool
	}{
		{
			testcase:          "Request with 10 usernames",
			url:               "/test/github",
			usernames:         []string{"nboy1", "bboy2", "aboy3", "aboy4", "eboy5", "fboy6", "gboy7", "hboy8", "iboy9", "jboy10"},
			expectedFirstName: "aboy3",
			result:            true,
			httpCode:          http.StatusOK,
		},
		{
			testcase:          "Request with 10 usernames - test cache",
			url:               "/test/github",
			usernames:         []string{"nboy1", "bboy2", "aboy3", "aboy4", "eboy5", "fboy6", "gboy7", "hboy8", "iboy9", "jboy10"},
			expectedFirstName: "aboy3",
			result:            true,
			shouldTestCache:   true,
			httpCode:          http.StatusOK,
		},
		{
			testcase:          "Request with more than 10 usernames",
			url:               "/test/github",
			usernames:         []string{"nboy1", "bboy2", "aboy3", "aboy4", "eboy5", "fboy6", "gboy7", "hboy8", "iboy9", "jboy10", "zboy"},
			expectedFirstName: "aboy3",
			message:           "request is invalid",
			httpCode:          http.StatusBadRequest,
		},
		{
			testcase:  "Request is empty",
			url:       "/test/github",
			usernames: []string{},
			message:   "request is invalid",
			httpCode:  http.StatusBadRequest,
		},
		{
			testcase:          "Has Github client error",
			url:               "/test/github",
			usernames:         []string{"boyHasClientError"},
			expectedFirstName: "",
			hasClientError:    true,
			result:            true,
			httpCode:          http.StatusOK,
		},
		{
			testcase:                  "Should error parsing request",
			url:                       "/test/github",
			shouldErrorParsingRequest: true,
			message:                   "Failed to parse request.",
			httpCode:                  http.StatusBadRequest,
		},
		{
			testcase:  "Source is not github",
			url:       "/test/bitbucket",
			usernames: []string{"nboy1", "bboy2", "aboy3", "aboy4", "eboy5", "fboy6", "gboy7", "hboy8", "iboy9", "jboy10"},
			message:   "Source not found.",
			httpCode:  http.StatusNotFound,
		},
	}

	for _, tc := range scenarios {
		t.Run(tc.testcase, func(t *testing.T) {
			util.GetCache().Clear()
			mockGithubClient := mock.Builder().
				GetGithubUserHasError(tc.hasClientError).
				Build()

			factory := service.NewAccountDetailServiceFactory(mockGithubClient)
			controller := New(factory)

			router := gin.Default()
			router.POST("/test/:source", controller.GetUserAccountDetails)

			var reqBytes []byte
			if !tc.shouldErrorParsingRequest {
				req := GetAccountDetailsRequest{Users: tc.usernames}
				r, err := json.Marshal(req)
				g.Expect(err).To(BeNil())
				reqBytes = r
			}

			req, err := http.NewRequest("POST", tc.url, bytes.NewReader(reqBytes))
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

			if tc.result && !tc.hasClientError {
				sampleResponse := response.UserDetails[0]
				g.Expect(sampleResponse.Name).To(Equal(tc.expectedFirstName))
			}
		})
	}
}
