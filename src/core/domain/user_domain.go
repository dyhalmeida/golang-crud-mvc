package domain

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomain struct {
	Email string
	Password string
	Name string
	Age int8
}

func NewUserDomain(name, email, password string, age int8) UserDomainInterface {
	return &UserDomain{
		Name: name,
		Email: email,
		Password: password,
		Age: age,
	}
}


func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
