package model

type GetGithubUserDetails struct {
	Name        string
	Login       string
	Company     string
	Followers   int
	PublicRepos int
}
