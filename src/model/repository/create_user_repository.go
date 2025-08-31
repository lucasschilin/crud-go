package repository

import (
	"github.com/lucasschilin/crud-go/src/configuration/rest_err"
	"github.com/lucasschilin/crud-go/src/model"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	//
	return nil, nil
}
