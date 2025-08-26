package model

import (
	"fmt"

	"github.com/lucasschilin/crud-go/src/configuration/logger"
	"github.com/lucasschilin/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) UpdateUser(string) *rest_err.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "update-user"))

	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
