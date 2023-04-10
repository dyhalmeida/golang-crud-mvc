package usercontroller

import (
	"fmt"

	resterrors "github.com/dyhalmeida/golang-crud-mvc/src/configuration/restErrors"
	"github.com/dyhalmeida/golang-crud-mvc/src/controller/model/request"
	"github.com/dyhalmeida/golang-crud-mvc/src/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	
	if utils.HasError(err) {
		restErr := resterrors.NewBadRequestError(fmt.Sprintf("There are some incorrect field, error=%s", err.Error()))
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