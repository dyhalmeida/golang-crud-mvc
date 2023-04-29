package usecase

import (
	resterrors "github.com/dyhalmeida/golang-crud-mvc/src/configuration/restErrors"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/domain"
)

type UserUsecaseInterface interface {
	CreateUser(domain.UserDomainInterface) (domain.UserDomainInterface, *resterrors.Error)
	UpdateUser(string, domain.UserDomainInterface) *resterrors.Error
	ShowUser(string) (domain.UserDomainInterface, *resterrors.Error)
	DeleteUser(string) *resterrors.Error
}
