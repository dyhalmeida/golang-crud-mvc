package converter

import (
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/domain"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/entity"
)

func EntityToDomain(userEntity entity.UserEntity) domain.UserDomainInterface {
	userDomain := domain.NewUserDomain(
		userEntity.Name,
		userEntity.Email,
		userEntity.Password,
		userEntity.Age,
	)
	userDomain.SetId(userEntity.Id.Hex())
	return userDomain
}