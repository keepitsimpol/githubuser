package port

import (
	"context"

	"github.com/keepitsimpol/githubuser/internal/core/constant/errorcode"
	"github.com/keepitsimpol/githubuser/internal/core/model"
)

type AccountDetailService interface {
	GetAccountDetails(users model.GetAccountDetailRequest, ctx context.Context) (response []interface{}, appError errorcode.AppErrorCode, err error)
}
