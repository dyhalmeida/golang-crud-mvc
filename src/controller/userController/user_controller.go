package usercontroller

import (
	"net/http"

	"github.com/dyhalmeida/golang-crud-mvc/src/configuration/logger"
	resterrors "github.com/dyhalmeida/golang-crud-mvc/src/configuration/restErrors"
	"github.com/dyhalmeida/golang-crud-mvc/src/configuration/validation"
	"github.com/dyhalmeida/golang-crud-mvc/src/controller/model/request"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/domain"
	"github.com/dyhalmeida/golang-crud-mvc/src/model/core/usecase"
	"github.com/dyhalmeida/golang-crud-mvc/src/utils"
	"github.com/dyhalmeida/golang-crud-mvc/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	userCreated, err := uc.userUsecase.CreateUser(domain.NewUserDomain(
		userRequest.Email,
		userRequest.Email,
		userRequest.Password,
		userRequest.Age,
	))
	if utils.HasError(err) {
		logger.Error("Error trying to call userUsecase.CreateUser", err, zap.String("flow", "userController.CreateUser"))
		context.JSON(err.Code, err)
		return
	}
	
	logger.Info(
		"User created successfully in user_controller",
		zap.String("userId", userCreated.GetId()),
		zap.String("flow", "CreateUser in user_controller"),
	)
	context.JSON(http.StatusOK, view.ConvertDomainToResponse(userCreated))

}

func (uc *userController) ShowUser(ctx *gin.Context) {
	logger.Info(
		"Init userController.ShowUser",
		zap.String("flow", "ShowUser"),
	)
	userId := ctx.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); utils.HasError(err) {
		logger.Error(
			"Error trying to validate userId",
			err,
			zap.String("flow", "userController.ShowUser"),
		)
		errorMessage := resterrors.NewBadRequestError("Invalid userId")
		ctx.JSON(errorMessage.Code, errorMessage)
	}

	userDomain, err := uc.userUsecase.ShowUser(userId)
	if utils.HasError(err) {
		logger.Error(
			"Error trying to call userUsecase.ShowUser",
			err,
			zap.String("flow", "userController.ShowUser"),
		)
		ctx.JSON(err.Code, err)
		return
	}
	logger.Info(
		"userController.ShowUser executed with success",
		zap.String("flow", "ShowUser"),
	)
	ctx.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userController) UpdateUser(context *gin.Context) {

}

func (uc *userController) DeleteUser(context *gin.Context) {

}