package domain

import resterrors "github.com/dyhalmeida/golang-crud-mvc/src/configuration/restErrors"

func (*UserDomain) DeleteUser(string) *resterrors.Error {
	return nil
}