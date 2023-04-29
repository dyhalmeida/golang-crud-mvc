package repository

import (
	resterrors "github.com/dyhalmeida/golang-crud-mvc/src/configuration/restErrors"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/domain"
)

type UserRepositoryInterface interface {
	CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *resterrors.Error)
}