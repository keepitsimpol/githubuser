package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/keepitsimpol/githubuser/internal/core/constant"
	"github.com/sirupsen/logrus"
)

func AppendRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := logrus.New().WithField(constant.RequestID, uuid.NewString())
		ctx.Set(constant.Logger, log)
	}
}
