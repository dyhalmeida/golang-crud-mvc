package view

import (
	"github.com/dyhalmeida/golang-crud-mvc/src/controller/model/response"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/domain"
)

func ConvertDomainToResponse(userDomain domain.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID: userDomain.GetId(),
		Email: userDomain.GetEmail(),
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(),
	}
}