package view

import (
	"github.com/lucasschilin/crud-go/src/controller/model/response"
	"github.com/lucasschilin/crud-go/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.User {
	return response.User{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
