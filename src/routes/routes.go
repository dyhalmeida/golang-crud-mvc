package routes

import (
	usercontroller "github.com/dyhalmeida/golang-crud-mvc/src/controller/userController"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.POST("/user", usercontroller.Create)
	r.GET("/user/:email", usercontroller.Show)
	r.PUT("/user/:id", usercontroller.Update)
	r.DELETE("/user/:id", usercontroller.Delete)

}