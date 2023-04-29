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
		"Init userRepository.CreateUser",
		zap.String("flow", "CreatseUser"),
	)
	userCollectionName := os.Getenv(MONGO_USER_COLLECTION)
	collection := userRepository.dbConnection.Collection(userCollectionName)
	userEntity := converter.DomainToEntity(userDomain)
	result, err := collection.InsertOne(context.Background(), userEntity)
	if utils.HasError(err) {
		logger.Error(
			"Error trying to collection.InsertOne in userRepository.CreateUser",
			err,
			zap.String("flow", "CreateUser"),
		)
		return nil, resterrors.NewInternalServerError(err.Error())
	}
	userEntity.Id = result.InsertedID.(primitive.ObjectID)
	logger.Info(
		"userRepository.CreateUser executed with success",
		zap.String("userId", userEntity.Id.Hex()),
		zap.String("flow", "CreateUser"),
	)
	return converter.EntityToDomain(*userEntity), nil
}

func (userRepository *userRepository) ShowUser(userID string) (domain.UserDomainInterface, *resterrors.Error) {
	logger.Info(
		"Init ShowUser in userRepository",
		zap.String("flow", "ShowUser"),
	)
	userCollectionName := os.Getenv(MONGO_USER_COLLECTION)
	collection := userRepository.dbConnection.Collection(userCollectionName)
	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.D{{ Key: "_id", Value: objectId }}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if utils.HasError(err) && err == mongo.ErrNoDocuments {
		errorMessage := fmt.Sprintf("User not found with this ID: %s", userID)
		logger.Error(
			errorMessage,
			err,
			zap.String("flow", "ShowUser"),
		)
		return nil, resterrors.NewNotFoundError(errorMessage)
	}
	if utils.HasError(err) {
		errorMessage := "Error trying to find user by ID"
		logger.Error(
			errorMessage,
			err,
			zap.String("flow", "ShowUser"),
		)
		return nil, resterrors.NewInternalServerError(errorMessage)
	}
	logger.Info(
		"ShowUser in userRepository executed with success",
		zap.String("userId", userEntity.Id.Hex()),
		zap.String("email", userEntity.Email),
		zap.String("flow", "ShowUser"),
	)
	return converter.EntityToDomain(*userEntity), nil
}