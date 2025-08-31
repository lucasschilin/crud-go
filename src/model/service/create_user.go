package service

import (
	"github.com/lucasschilin/crud-go/src/configuration/logger"
	"github.com/lucasschilin/crud-go/src/configuration/rest_err"
	"github.com/lucasschilin/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	journeyTag := zap.String("journey", "create-user")
	logger.Info("Init createUser model", journeyTag)

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, journeyTag)
		return nil, err
	}

	logger.Info(
		"CreateUser controller executed successfully",
		zap.String("userID", userDomainRepository.GetID()),
		journeyTag,
	)
	return userDomainRepository, nil
}
