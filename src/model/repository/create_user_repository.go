package repository

import (
	"context"
	"os"

	"github.com/lucasschilin/crud-go/src/configuration/logger"
	"github.com/lucasschilin/crud-go/src/configuration/rest_err"
	"github.com/lucasschilin/crud-go/src/model"
	"github.com/lucasschilin/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	journeyTag := zap.String("journey", "create-user")

	logger.Info("Init createUser repository", journeyTag)

	collection := ur.databaseConnection.Collection(os.Getenv(COLLECTION_NAME))

	userRepository := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), userRepository)
	if err != nil {
		logger.Error("Error trying to insert in collection", err, journeyTag)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userRepository.ID = result.InsertedID.(bson.ObjectID)

	logger.Info(
		"CreateUser repository executed successfully",
		zap.String("userID", userRepository.ID.Hex()),
		journeyTag,
	)
	return converter.ConvertEntityToDomain(*userRepository), nil
}
