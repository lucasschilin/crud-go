package model

import (
	"github.com/lucasschilin/crud-go/src/configuration/logger"
	"github.com/lucasschilin/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser() model", zap.String("journey", "create-user"))

	ud.EncryptPassword()

	return nil
}
