package presentation

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keepitsimpol/githubuser/internal/core/constant/errorcode"
	"github.com/keepitsimpol/githubuser/internal/core/model"
	"github.com/keepitsimpol/githubuser/internal/core/port"
	"github.com/keepitsimpol/githubuser/internal/core/util"
)

type accountDetailController struct {
	service port.AccountDetailService
}

func New(service port.AccountDetailService) *accountDetailController {
	controller := new(accountDetailController)
	controller.service = service
	return controller
}

type GetAccountDetailsRequest struct {
	Users []string `json:"users"`
}

type GetAccountDetailsResponse struct {
	Result      bool   `json:"result"`
	Message     string `json:"message"`
	UserDetails []interface{}
}

// GetUserAccountDetails godoc
// @Summary      Get details of all provided github users
// @Description  Get details of all provided github users
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        account  body      GetAccountDetailsRequest  true  "List of users"
// @Success      200      {object}  GetAccountDetailsResponse
// @Failure      400      {object}  GetAccountDetailsResponse
// @Failure      404      {object}  GetAccountDetailsResponse
// @Failure      500      {object}  GetAccountDetailsResponse
// @Router       /users/github [post]
func (c accountDetailController) GetUserAccountDetails(ctx *gin.Context) {
	util.Infoln(ctx, "Start Getting all users account details")

	var request GetAccountDetailsRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		util.Errorf(ctx, "Failed to parse request with error: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, GetAccountDetailsResponse{Message: "Failed to parse request."})
		return
	}

	response, errorType, err := c.service.GetAccountDetails(model.GetAccountDetailRequest{UserNames: request.Users}, ctx)
	httpCode := c.convertAppErrorCodeToHttpCode(errorType, ctx)
	if err != nil {
		ctx.JSON(httpCode, GetAccountDetailsResponse{Message: err.Error()})
	}

	ctx.JSON(httpCode, GetAccountDetailsResponse{
		Result:      true,
		UserDetails: response,
	})
	util.Infoln(ctx, "End Getting all github users detail")
}

func (c accountDetailController) convertAppErrorCodeToHttpCode(errorType errorcode.AppErrorCode, ctx context.Context) int {
	util.Infof(ctx, "Converting app error: %d to http code", errorType)
	if errorType == errorcode.NoError {
		return http.StatusOK
	} else if errorType == errorcode.InvalidRequest {
		return http.StatusBadRequest
	} else if errorType == errorcode.InternalError {
		return http.StatusInternalServerError
	} else if errorType == errorcode.DependencyError {
		return http.StatusFailedDependency
	} else {
		util.Errorf(ctx, "Unsupported error type: %d", errorType)
		return http.StatusInternalServerError
	}
}
