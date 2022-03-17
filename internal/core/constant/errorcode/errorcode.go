package errorcode

type AppErrorCode int

const (
	NoError = iota
	InvalidRequest
	InternalError
	DependencyError
)
