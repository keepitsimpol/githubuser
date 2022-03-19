package port

import "context"

type GithubClient interface {
	GetGithubUser(user string, ctx context.Context) (GetGithubUserResponse, error)
}

type GetGithubUserResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Login       string `json:"login"`
	Company     string `json:"company"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	PublicRepos int    `json:"public_repos"`
	Type        string `json:"type"`
	Blog        string `json:"blog"`
	Email       string `json:"email"`
}
