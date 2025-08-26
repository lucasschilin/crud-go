package model

import (
	"github.com/lucasschilin/crud-go/src/configuration/logger"
	"github.com/lucasschilin/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) DeleteUser() *rest_err.RestErr {
	logger.Info("Init deleteUser model", zap.String("journey", "delete-user"))

	return nil
}
