package domain

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"github.com/dyhalmeida/golang-crud-mvc/src/utils"
)

type userDomain struct {
	id string
	email string
	password string
	name string
	age int8
}

func NewUserDomain(name, email, password string, age int8) UserDomainInterface {
	return &userDomain{
		name: name,
		email: email,
		password: password,
		age: age,
	}
}
func(ud *userDomain) GetId() string {
	return ud.id
}

func(ud *userDomain) GetEmail() string {
	return ud.email
}

func(ud *userDomain) GetPassword() string {
	return ud.password
}

func(ud *userDomain) GetName() string {
	return ud.name
}

func(ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) SetId(id string) {
	ud.id = id
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) ToString() (string, error) {
	json, err := json.Marshal(ud)
	if utils.HasError(err) {
		return "", err
	}
	return string(json), nil
}
