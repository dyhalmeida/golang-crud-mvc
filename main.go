package main

import (
	"context"
	"log"

	"github.com/dyhalmeida/golang-crud-mvc/src/configuration/database/mongodb"
	"github.com/dyhalmeida/golang-crud-mvc/src/configuration/logger"
	usercontroller "github.com/dyhalmeida/golang-crud-mvc/src/controller/userController"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/repository"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/usecase"
	"github.com/dyhalmeida/golang-crud-mvc/src/routes"
	"github.com/dyhalmeida/golang-crud-mvc/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initDependencies(database *mongo.Database) usercontroller.UserControllerInterface {
	userRepository := repository.NewUserRepository(database)
	userUsecase := usecase.NewUserUsecase(userRepository)
	return usercontroller.NewUserController(userUsecase)
}

func main()  {

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if utils.HasError(err) {
		logger.Error("Error trying to connect to database in main", err, zap.String("flow", "main"))
		log.Fatal("Error trying to connect to database")
		return
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, initDependencies(database))
	
	if err := router.Run(":3333"); err != nil {
		log.Fatal(err)
	}
}