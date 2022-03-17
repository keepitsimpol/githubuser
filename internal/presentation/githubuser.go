package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keepitsimpol/githubuser/internal/core/constant"
	"github.com/keepitsimpol/githubuser/internal/core/constant/errorcode"
	"github.com/keepitsimpol/githubuser/internal/core/gateway"
	"github.com/keepitsimpol/githubuser/internal/core/model"
	"github.com/keepitsimpol/githubuser/internal/core/service"
	"github.com/sirupsen/logrus"
)

type githubUserController struct {
	githubClient gateway.GithubClient
}

func New(githubClient gateway.GithubClient) *githubUserController {
	controller := new(githubUserController)
	controller.githubClient = githubClient
	return controller
}

type GetGithuUsersRequest struct {
	Users []string `json:"users"`
}

type GetGithubUsersResponse struct {
	Result      bool   `json:"result"`
	Message     string `json:"message"`
	UserDetails []model.GetGithubUserDetails
}

// GetGithubUsers godoc
// @Summary      Get details of all provided github users
// @Description  Get details of all provided github users
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        account  body      GetGithuUsersRequest  true  "List of users"
// @Success      200      {object}  GetGithubUsersResponse
// @Failure      400      {object}  GetGithubUsersResponse
// @Failure      404      {object}  GetGithubUsersResponse
// @Failure      500      {object}  GetGithubUsersResponse
// @Router       /users [post]
func (c githubUserController) GetGithubUsers(ctx *gin.Context) {
	logger := ctx.Value(constant.Logger).(*logrus.Entry)
	logger.Println("Start Getting all github users detail")

	var request GetGithuUsersRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		logger.Errorf("Failed to parse request with error: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, GetGithubUsersResponse{Message: "Failed to parse request."})
		return
	}

	service := service.New(logger)
	response, errorType, err := service.GetGithubUsers(request.Users)
	httpCode := c.convertAppErrorCodeToHttpCode(errorType, logger)
	if err != nil {
		ctx.JSON(httpCode, GetGithubUsersResponse{Message: err.Error()})
	}

	ctx.JSON(httpCode, GetGithubUsersResponse{
		Result:      true,
		UserDetails: response,
	})
	logger.Println("End Getting all github users detail")
}

func (c githubUserController) convertAppErrorCodeToHttpCode(errorType errorcode.AppErrorCode, logger *logrus.Entry) int {
	logger.Infof("Converting app error: %d to http code", errorType)
	if errorType == errorcode.NoError {
		return http.StatusOK
	} else if errorType == errorcode.InvalidRequest {
		return http.StatusBadRequest
	} else if errorType == errorcode.InternalError {
		return http.StatusInternalServerError
	} else if errorType == errorcode.DependencyError {
		return http.StatusFailedDependency
	} else {
		logrus.Errorf("Unsupported error type: %d", errorType)
		return http.StatusInternalServerError
	}
}
