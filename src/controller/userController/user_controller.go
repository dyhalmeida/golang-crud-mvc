package usercontroller

import (
	"fmt"

	"github.com/dyhalmeida/golang-crud-mvc/src/configuration/validation"
	"github.com/dyhalmeida/golang-crud-mvc/src/controller/model/request"
	"github.com/dyhalmeida/golang-crud-mvc/src/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	var userRequest request.UserRequest

	err := context.ShouldBindJSON(&userRequest)
	
	if utils.HasError(err) {
		restErr := validation.ValidateError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)

}

func ShowUser(context *gin.Context) {
}

func UpdateUser(context *gin.Context) {

}

func DeleteUser(context *gin.Context) {

}