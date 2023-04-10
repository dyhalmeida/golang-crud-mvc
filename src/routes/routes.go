package routes

import (
	usercontroller "github.com/dyhalmeida/golang-crud-mvc/src/controller/userController"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.POST("/user", usercontroller.CreateUser)
	r.GET("/user/:email", usercontroller.ShowUser)
	r.PUT("/user/:email", usercontroller.UpdateUser)
	r.DELETE("/user/:email", usercontroller.DeleteUser)

}