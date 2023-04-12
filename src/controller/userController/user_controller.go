package usercontroller

import (
	"fmt"

	"github.com/dyhalmeida/golang-crud-mvc/src/configuration/validation"
	"github.com/dyhalmeida/golang-crud-mvc/src/controller/model/request"
	"github.com/dyhalmeida/golang-crud-mvc/src/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	
	if utils.HasError(err) {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)

}

func ShowUser(c *gin.Context) {
}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}