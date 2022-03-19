package port

type AccountDetailFactory interface {
	GetAccountDetailService(source string) (AccountDetailService, error)
}
