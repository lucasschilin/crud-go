package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/lucasschilin/crud-go/src/configuration/logger"
	"github.com/lucasschilin/crud-go/src/configuration/rest_err"
	"github.com/lucasschilin/crud-go/src/model"
	"github.com/lucasschilin/crud-go/src/model/repository/entity"
	"github.com/lucasschilin/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	journeyTag := zap.String("journey", "find-user-by-email")

	logger.Info("Init findUserByEmail repository", journeyTag)

	collection := ur.databaseConnection.Collection(os.Getenv(COLLECTION_NAME))

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", email,
			)
			logger.Error(errorMessage, err, journeyTag)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, journeyTag)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info(
		"FinUserByEmail repository executed successfully",
		zap.String("email", email),
		zap.String("userId", string(userEntity.ID.Hex())),
		journeyTag,
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	journeyTag := zap.String("journey", "find-user-by-id")

	logger.Info("Init findUserByID repository", journeyTag)

	collection := ur.databaseConnection.Collection(os.Getenv(COLLECTION_NAME))

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this ID: %s", id,
			)
			logger.Error(errorMessage, err, journeyTag)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by ID"
		logger.Error(errorMessage, err, journeyTag)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info(
		"FinUserByID repository executed successfully",
		zap.String("userId", string(userEntity.ID.Hex())),
		journeyTag,
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}
