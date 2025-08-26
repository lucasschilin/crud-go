package model

import (
	"github.com/lucasschilin/crud-go/src/configuration/logger"
	"github.com/lucasschilin/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (u *userDomain) FindUser(string) (*userDomain, *rest_err.RestErr) {
	logger.Info("Init findUser model", zap.String("journey", "find-user"))

	return nil, nil
}
