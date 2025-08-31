package repository

import (
	"context"
	"os"

	"github.com/lucasschilin/crud-go/src/configuration/logger"
	"github.com/lucasschilin/crud-go/src/configuration/rest_err"
	"github.com/lucasschilin/crud-go/src/model"
	"github.com/lucasschilin/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	COLLECTION_NAME = "DATABASE_USER_COLLECTION_NAME"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser repository")

	collection := ur.databaseConnection.Collection(os.Getenv(COLLECTION_NAME))

	userRepository := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), userRepository)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userRepository.ID = result.InsertedID.(bson.ObjectID)

	return converter.ConvertEntityToDomain(*userRepository), nil
}
