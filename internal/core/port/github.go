package port

import "context"

type GithubClient interface {
	GetGithubUser(user string, ctx context.Context) (GetGithubUserResponse, error)
}

type GetGithubUserResponse struct {
	Name        string `json:"name"`
	Login       string `json:"login"`
	Company     string `json:"company"`
	Followers   int    `json:"followers"`
	PublicRepos int    `json:"public_repos"`
}
