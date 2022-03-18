package model

type GetAccountDetailRequest struct {
	UserNames []string `validate:"min=1,max=10"`
}

type GetAccountDetailResponse struct {
	Name        string
	Login       string
	Company     string
	Followers   int
	PublicRepos int
}
