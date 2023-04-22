package usercontroller

import (
	"net/http"

	"github.com/dyhalmeida/golang-crud-mvc/src/configuration/logger"
	"github.com/dyhalmeida/golang-crud-mvc/src/configuration/validation"
	"github.com/dyhalmeida/golang-crud-mvc/src/controller/model/request"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/domain"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/usecase"
	"github.com/dyhalmeida/golang-crud-mvc/src/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface domain.UserDomainInterface
)

type UserControllerInterface interface {
	CreateUser(context *gin.Context)
	ShowUser(context *gin.Context)
	UpdateUser(context *gin.Context)
	DeleteUser(context *gin.Context)
}

type userController struct {
	usecase usecase.UserUsecaseInterface
}

func CreateUser(context *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("flow", "CreateUser"))
	var userRequest request.UserRequest

	err := context.ShouldBindJSON(&userRequest)
	
	if utils.HasError(err) {
		logger.Error("Error trying to validate user data", err, zap.String("flow", "CreateUser"))
		restErr := validation.ValidateError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	domain := domain.NewUserDomain(
		userRequest.Email,
		userRequest.Email,
		userRequest.Password,
		userRequest.Age,
	)
	
	usecase := usecase.NewUserUsecase()

	if err := usecase.CreateUser(domain); err != nil {
		context.JSON(err.Code, err)
		return
	}

	context.String(http.StatusOK, "")
	logger.Info("User created successfully", zap.String("flow", "CreateUser"))

}

func ShowUser(context *gin.Context) {
}

func UpdateUser(context *gin.Context) {

}

func DeleteUser(context *gin.Context) {

}