package usecase

import (
	"fmt"

	resterrors "github.com/dyhalmeida/golang-crud-mvc/src/configuration/restErrors"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/domain"
)

type userUsecase struct {}

func NewUserUsecase() UserUsecaseInterface {
	return &userUsecase{}
}

func (ud *userUsecase) CreateUser(
	userDomain domain.UserDomainInterface,
) *resterrors.Error {
	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetName())
	return nil
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