package domain

type UserDomainInterface interface {
	GetId() string
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	SetId(string)
	ToString() (string, error)
	EncryptPassword()
}