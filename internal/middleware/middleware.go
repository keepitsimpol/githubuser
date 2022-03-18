package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/keepitsimpol/githubuser/internal/core/constant"
)

func AppendRequestID(ctx *gin.Context) {
	ctx.Set(constant.RequestID, uuid.NewString())
	ctx.Next()
}
