package main

import (
	"log"

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
	routes.InitRoutes(&router.RouterGroup)
	
	if err := router.Run(":3333"); err != nil {
		log.Fatal(err)
	}
}