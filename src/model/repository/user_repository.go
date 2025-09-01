package repository

import (
	"github.com/lucasschilin/crud-go/src/configuration/rest_err"
	"github.com/lucasschilin/crud-go/src/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	COLLECTION_NAME = "DATABASE_USER_COLLECTION_NAME"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		databaseConnection: database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
}
