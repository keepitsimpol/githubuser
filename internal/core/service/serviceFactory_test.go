package service

import (
	"testing"

	"github.com/keepitsimpol/githubuser/internal/core/mock"
	. "github.com/onsi/gomega"
)

func TestGetAccountDetailService(t *testing.T) {
	g := NewGomegaWithT(t)
	scenarios := []struct {
		testcase string
		source   string
		hasError bool
	}{
		{
			testcase: "Get service for github",
			source:   "github",
		},
		{
			testcase: "Get service for bitbucket",
			source:   "bitbucket",
			hasError: true,
		},
	}

	for _, tc := range scenarios {
		t.Run(tc.testcase, func(t *testing.T) {
			mockGithubClient := mock.Builder().Build()
			factory := NewAccountDetailServiceFactory(mockGithubClient)

			service, err := factory.GetAccountDetailService(tc.source)
			if tc.hasError {
				g.Expect(err).ToNot(BeNil())
				g.Expect(service).To(BeNil())
			} else {
				g.Expect(err).To(BeNil())
				g.Expect(service).ToNot(BeNil())
			}
		})
	}
}
