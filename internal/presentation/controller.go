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
	serviceFactory port.AccountDetailFactory
}

func New(serviceFactory port.AccountDetailFactory) *accountDetailController {
	service := new(accountDetailController)
	service.serviceFactory = serviceFactory
	return service
}

type GetAccountDetailsRequest struct {
	Users []string `json:"users"`
}

type GetAccountDetailsResponse struct {
	Result      bool          `json:"result"`
	Message     string        `json:"message"`
	UserDetails []UserDetails `json:"userDetails"`
}

type UserDetails struct {
	Name        string `json:"name,omitempty"`
	Login       string `json:"login,omitempty"`
	Company     string `json:"company,omitempty"`
	Followers   int    `json:"followers"`
	PublicRepos int    `json:"publicRepos,omitempty"`
}

// GetUserAccountDetails godoc
// @Summary      Get details of all provided github users
// @Description  Get details of all provided github users
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        account  body      GetAccountDetailsRequest  true  "List of users"
// @Param        source   path      string                    true  "Account source"
// @Success      200      {object}  GetAccountDetailsResponse
// @Failure      400      {object}  GetAccountDetailsResponse
// @Failure      404      {object}  GetAccountDetailsResponse
// @Failure      500      {object}  GetAccountDetailsResponse
// @Router       /users/{source} [post]
func (c accountDetailController) GetUserAccountDetails(ctx *gin.Context) {
	util.Infoln(ctx, "Start Getting all users account details")

	var request GetAccountDetailsRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		util.Errorf(ctx, "Failed to parse request with error: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, GetAccountDetailsResponse{Message: "Failed to parse request."})
		return
	}

	service, err := c.serviceFactory.GetAccountDetailService(ctx.Param("source"))
	if err != nil {
		util.Errorf(ctx, "Error in getting service: %s", err.Error())
		ctx.JSON(http.StatusNotFound, GetAccountDetailsResponse{Message: "Source not found."})
		return
	}

	response, errorType, err := service.GetAccountDetails(model.GetAccountDetailRequest{UserNames: request.Users}, ctx)
	httpCode := c.convertAppErrorCodeToHttpCode(errorType, ctx)
	if err != nil {
		ctx.JSON(httpCode, GetAccountDetailsResponse{Message: err.Error()})
		return
	}

	ctx.JSON(httpCode, GetAccountDetailsResponse{
		Result:      true,
		UserDetails: c.convertToUserDetails(response),
	})
	util.Infoln(ctx, "End Getting all github users detail")
}

func (c accountDetailController) convertAppErrorCodeToHttpCode(errorType errorcode.AppErrorCode, ctx context.Context) int {
	util.Infof(ctx, "Converting app error: %d to http code", errorType)
	if errorType == errorcode.NoError {
		return http.StatusOK
	} else if errorType == errorcode.InvalidRequest {
		return http.StatusBadRequest
	} else {
		util.Errorf(ctx, "Unsupported error type: %d", errorType)
		return http.StatusInternalServerError
	}
}

func (c accountDetailController) convertToUserDetails(response []model.GetAccountDetailResponse) (userDetails []UserDetails) {
	for _, accountDetail := range response {
		userDetails = append(userDetails, UserDetails{
			Name:        accountDetail.Name,
			Login:       accountDetail.Login,
			Company:     accountDetail.Company,
			Followers:   accountDetail.Followers,
			PublicRepos: accountDetail.PublicRepos,
		})
	}
	return userDetails
}
