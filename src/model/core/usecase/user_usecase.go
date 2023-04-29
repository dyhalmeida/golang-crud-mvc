package usecase

import (
	"github.com/dyhalmeida/golang-crud-mvc/src/configuration/logger"
	resterrors "github.com/dyhalmeida/golang-crud-mvc/src/configuration/restErrors"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/domain"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/repository"
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
		"Init CreatseUser in userUsecase in user_usecase",
		zap.String("flow", "CreatseUser"),
	)
	userDomain.EncryptPassword()
	userCreated, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error(
			"Error trying to call userRepository.CreateUser in user_usecase",
			err,
			zap.String("flow", "CreateUser"),
		)
		return nil, err
	}
	logger.Info(
		"CreateUser userUsecase executed successfully in user_usecase",
		zap.String("userId", userCreated.GetId()),
		zap.String("flow", "CreateUser"),
	)
	return userCreated, nil
}

func (*userUsecase) DeleteUser(userID string) *resterrors.Error {
	return nil
}

func (*userUsecase) ShowUser(userID string) (*domain.UserDomainInterface, *resterrors.Error) {
	return nil, nil
}

func (*userUsecase) UpdateUser(
	userID string, 
	userDomain domain.UserDomainInterface,
) *resterrors.Error {
	return nil
}