package usercontroller

import (
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
	logger.Info("Init CreateUser controller", zap.String("flow", "CreateUser"))
	var userRequest request.UserRequest

	err := context.ShouldBindJSON(&userRequest)
	
	if utils.HasError(err) {
		logger.Error("Error trying to validate user data", err, zap.String("flow", "CreateUser"))
		restErr := validation.ValidateError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	userDomain := domain.NewUserDomain(
		userRequest.Email,
		userRequest.Email,
		userRequest.Password,
		userRequest.Age,
	)

	if err := uc.userUsecase.CreateUser(userDomain); err != nil {
		context.JSON(err.Code, err)
		return
	}

	context.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
	logger.Info("User created successfully", zap.String("flow", "CreateUser"))

}

func (uc *userController) ShowUser(context *gin.Context) {
}

func (uc *userController) UpdateUser(context *gin.Context) {

}

func (uc *userController) DeleteUser(context *gin.Context) {

}