package model

type GetAccountDetailRequest struct {
	UserNames []string `validate:"min=1,max=10"`
}
