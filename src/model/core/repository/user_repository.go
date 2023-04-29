package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/dyhalmeida/golang-crud-mvc/src/configuration/logger"
	resterrors "github.com/dyhalmeida/golang-crud-mvc/src/configuration/restErrors"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/domain"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/entity"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/entity/converter"
	"github.com/dyhalmeida/golang-crud-mvc/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)


var (
	MONGO_USER_COLLECTION = "MONGO_USER_COLLECTION"
)

type userRepository struct {
	dbConnection *mongo.Database
}

func NewUserRepository(database *mongo.Database) UserRepositoryInterface {
	return &userRepository{
		dbConnection: database,
	}
}

func (userRepository *userRepository) CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *resterrors.Error) {
	logger.Info(
		"Init CreateUser in userRepository",
		zap.String("flow", "CreatseUser"),
	)
	userCollectionName := os.Getenv(MONGO_USER_COLLECTION)
	collection := userRepository.dbConnection.Collection(userCollectionName)
	userEntity := converter.DomainToEntity(userDomain)
	result, err := collection.InsertOne(context.Background(), userEntity)
	if utils.HasError(err) {
		logger.Error(
			"Error trying to collection.InsertOne in user_repository",
			err,
			zap.String("flow", "CreateUser"),
		)
		return nil, resterrors.NewInternalServerError(err.Error())
	}
	userEntity.Id = result.InsertedID.(primitive.ObjectID)
	logger.Info(
		"User created successfully in user_repository",
		zap.String("userId", userEntity.Id.Hex()),
		zap.String("flow", "CreateUser"),
	)
	return converter.EntityToDomain(*userEntity), nil
}
