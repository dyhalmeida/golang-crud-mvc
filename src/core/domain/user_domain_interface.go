package domain

import resterrors "github.com/dyhalmeida/golang-crud-mvc/src/configuration/restErrors"

type UserDomainInterface interface {
	CreateUser() *resterrors.Error
	UpdateUser(string) *resterrors.Error
	ShowUser(string) (*UserDomain, *resterrors.Error)
	DeleteUser(string) *resterrors.Error
}
