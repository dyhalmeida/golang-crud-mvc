package domain

import (
	"fmt"

	resterrors "github.com/dyhalmeida/golang-crud-mvc/src/configuration/restErrors"
)

func (ud *UserDomain) CreateUser() *resterrors.Error {
	ud.EncryptPassword()
	fmt.Println(ud)
	return nil
}