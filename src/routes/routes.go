package routes

import (
	usercontroller "github.com/dyhalmeida/golang-crud-mvc/src/controller/userController"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController usercontroller.UserControllerInterface) {

	r.POST("/user", userController.CreateUser)
	r.GET("/user/:userId", userController.ShowUser)
	r.PUT("/user/:email", userController.UpdateUser)
	r.DELETE("/user/:email", userController.DeleteUser)

}