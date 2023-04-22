package usercontroller

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	CreateUser(context *gin.Context)
	ShowUser(context *gin.Context)
	UpdateUser(context *gin.Context)
	DeleteUser(context *gin.Context)
}
