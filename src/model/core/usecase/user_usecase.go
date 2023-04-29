package usecase

import (
	"github.com/dyhalmeida/golang-crud-mvc/src/configuration/logger"
	resterrors "github.com/dyhalmeida/golang-crud-mvc/src/configuration/restErrors"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/domain"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/repository"
	"github.com/dyhalmeida/golang-crud-mvc/src/utils"
	"go.uber.org/zap"
)

type userUsecase struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserUsecase(userRepository repository.UserRepositoryInterface) UserUsecaseInterface {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (ud *userUsecase) CreateUser(
	userDomain domain.UserDomainInterface,
) (domain.UserDomainInterface, *resterrors.Error) {
	logger.Info(
		"Init userUsecase.CreateUser",
		zap.String("flow", "CreatseUser"),
	)
	userDomain.EncryptPassword()
	userDomainCreated, err := ud.userRepository.CreateUser(userDomain)
	
	if utils.HasError(err) {
		logger.Error(
			"Error trying to call userRepository.CreateUser",
			err,
			zap.String("flow", "userUsecase.CreateUser"),
		)
		return nil, err
	}
	
	logger.Info(
		"userUsecase.CreateUser executed with success",
		zap.String("userId", userDomainCreated.GetId()),
		zap.String("flow", "CreateUser"),
	)
	return userDomainCreated, nil
}

func (*userUsecase) DeleteUser(userID string) *resterrors.Error {
	return nil
}

func (ud *userUsecase) ShowUser(userID string) (domain.UserDomainInterface, *resterrors.Error) {
	logger.Info(
		"Init UserUsecase.ShowUser",
		zap.String("flow", "ShowUser"),
	)
	userDomain, err := ud.userRepository.ShowUser(userID)
	if utils.HasError(err) {
		return nil, err
	}
	return userDomain, nil
}

func (*userUsecase) UpdateUser(
	userID string, 
	userDomain domain.UserDomainInterface,
) *resterrors.Error {
	return nil
}