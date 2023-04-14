package usercontroller

import (
	"net/http"

	"github.com/dyhalmeida/golang-crud-mvc/src/configuration/validation"
	"github.com/dyhalmeida/golang-crud-mvc/src/controller/model/request"
	"github.com/dyhalmeida/golang-crud-mvc/src/core/domain"
	"github.com/dyhalmeida/golang-crud-mvc/src/utils"
	"github.com/gin-gonic/gin"
)

var (
	UserDomainInterface domain.UserDomainInterface
)

func CreateUser(context *gin.Context) {
	var userRequest request.UserRequest

	err := context.ShouldBindJSON(&userRequest)
	
	if utils.HasError(err) {
		restErr := validation.ValidateError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	domain := domain.NewUserDomain(
		userRequest.Email,
		userRequest.Email,
		userRequest.Password,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		context.JSON(err.Code, err)
		return
	}

	context.String(http.StatusOK, "")

}

func ShowUser(context *gin.Context) {
}

func UpdateUser(context *gin.Context) {

}

func DeleteUser(context *gin.Context) {

}