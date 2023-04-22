package domain

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	EncryptPassword()
}