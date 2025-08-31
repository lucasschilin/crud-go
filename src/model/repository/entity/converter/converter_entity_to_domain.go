package converter

import (
	"github.com/lucasschilin/crud-go/src/model"
	"github.com/lucasschilin/crud-go/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	userDomain := model.NewUserDomain(
		entity.Email, entity.Password, entity.Name, entity.Age,
	)

	userDomain.SetID(entity.ID.Hex())

	return userDomain
}
