package converter

import (
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/domain"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/entity"
)

func DomainToEntity(userDomain domain.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email: userDomain.GetEmail(),
		Password: userDomain.GetPassword(),
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(),
	}
}