package usercontroller

import (
	"fmt"
	"net/http"

	"github.com/dyhalmeida/golang-crud-mvc/src/configuration/logger"
	"github.com/dyhalmeida/golang-crud-mvc/src/configuration/validation"
	"github.com/dyhalmeida/golang-crud-mvc/src/controller/model/request"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/domain"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/usecase"
	"github.com/dyhalmeida/golang-crud-mvc/src/utils"
	"github.com/dyhalmeida/golang-crud-mvc/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type userController struct {
	userUsecase usecase.UserUsecaseInterface
}

func NewUserController(userUseCase usecase.UserUsecaseInterface) UserControllerInterface {
	return &userController{
		userUsecase: userUseCase,
	}
}

func (uc *userController) CreateUser(context *gin.Context) {
	logger.Info("Init CreateUser controller in user_controller", zap.String("flow", "CreateUser in user_controller"))
	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); utils.HasError(err) {
		logger.Error("Error trying to validate user data in user_controller", err, zap.String("flow", "CreateUser in user_controller"))
		restErr := validation.ValidateError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	userCreated, _ := uc.userUsecase.CreateUser(domain.NewUserDomain(
		userRequest.Email,
		userRequest.Email,
		userRequest.Password,
		userRequest.Age,
	))

	fmt.Println(userCreated, "DIEGO ALMEIDA")
	
	// if utils.HasError(err) {
	// 	logger.Error("Error trying to call userUsecase.CreateUser in user_controller", err, zap.String("flow", "CreateUser in user_controller"))
	// 	context.JSON(err.Code, err)
	// 	return
	// }
	logger.Info(
		"User created successfully in user_controller",
		zap.String("userId", userCreated.GetId()),
		zap.String("flow", "CreateUser in user_controller"),
	)
	context.JSON(http.StatusOK, view.ConvertDomainToResponse(userCreated))

}

func (uc *userController) ShowUser(context *gin.Context) {
}

func (uc *userController) UpdateUser(context *gin.Context) {

}

func (uc *userController) DeleteUser(context *gin.Context) {

}