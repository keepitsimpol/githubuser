package util

import (
	"context"

	"github.com/keepitsimpol/githubuser/internal/core/constant"
	"github.com/sirupsen/logrus"
)

func Infoln(ctx context.Context, message string) {
	logger := logrus.WithField(constant.RequestID, ctx.Value(constant.RequestID))
	logger.Infoln(message)
}

func Infof(ctx context.Context, message string, args ...interface{}) {
	logger := logrus.WithField(constant.RequestID, ctx.Value(constant.RequestID))
	logger.Infof(message, args...)
}

func Errorln(message string, ctx context.Context) {
	logger := logrus.WithField(constant.RequestID, ctx.Value(constant.RequestID))
	logger.Errorln(message)
}

func Errorf(ctx context.Context, message string, args ...interface{}) {
	logger := logrus.WithField(constant.RequestID, ctx.Value(constant.RequestID))
	logger.Errorf(message, args...)
}
