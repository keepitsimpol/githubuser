package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type githubUserHandler struct{}

func New() *githubUserHandler {
	return new(githubUserHandler)
}

type GetGithubUserResponse struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

func (h githubUserHandler) GetGethubUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, GetGithubUserResponse{
			Result:  true,
			Message: "Hello world!",
		})
	}
}
