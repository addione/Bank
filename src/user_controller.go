package src

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/addione/New/manager"
	"github.com/addione/New/models"
	"github.com/addione/New/src/request"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userManager *manager.UserManager
}

func NewUserController(sdi *srcDiContainer) *userController {
	um := sdi.managerDIContainer.GetUserManager()
	return &userController{
		userManager: um,
	}
}

func (uc *userController) CreateUser(ctx *gin.Context) {
	createUserParams, err := request.ValidateCreateUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	u, err := uc.userManager.CreateNewUser(createUserParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, u)

}

func (uc *userController) GetUserById(context *gin.Context) {
	user, err := uc.getAndvalidateUserById(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": user})
}

func (uc *userController) UpdateUser(context *gin.Context) {
	LoggedInUserInfo, _ := context.Get("userInfo")

	fmt.Println(LoggedInUserInfo, "this is logged in user")

	user, err := uc.getAndvalidateUserById(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var userUpdateParams models.UserUpdateRequest
	err = context.ShouldBindJSON(&userUpdateParams)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = uc.userManager.UpdateUser(user.ID, &userUpdateParams)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "ok"})
	return
}

func (uc *userController) ListUsers(context *gin.Context) {
	context.JSON(http.StatusOK, uc.userManager.ListUsers())
}

func (uc *userController) getAndvalidateUserById(context *gin.Context) (*models.UserMysql, error) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		return nil, err
	}
	user, err := uc.userManager.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (uc *userController) Login(ctx *gin.Context) {
	var loginParams *models.UserLoginRequest
	err := ctx.BindJSON(&loginParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, err := uc.userManager.ValidateCredentialsAndGetToken(loginParams)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
