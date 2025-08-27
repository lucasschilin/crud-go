package service

import (
	"github.com/lucasschilin/crud-go/src/configuration/logger"
	"github.com/lucasschilin/crud-go/src/configuration/rest_err"
	"github.com/lucasschilin/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(
	userID string, userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "update-user"))

	return nil
}
