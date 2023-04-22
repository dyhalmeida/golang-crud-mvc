package main

import (
	"log"

	usercontroller "github.com/dyhalmeida/golang-crud-mvc/src/controller/userController"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/usecase"
	"github.com/dyhalmeida/golang-crud-mvc/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main()  {
	router := gin.Default()

	userUsecase := usecase.NewUserUsecase()
	userController := usercontroller.NewUserController(userUsecase)

	routes.InitRoutes(&router.RouterGroup, userController)
	
	if err := router.Run(":3333"); err != nil {
		log.Fatal(err)
	}
}