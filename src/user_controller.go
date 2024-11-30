package src

import (
	"net/http"
	"strconv"

	"github.com/addione/New/manager"
	"github.com/addione/New/models"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userManager *manager.UserManager
}

func NewUserController() *userController {
	um := manager.NewManagerDIContainer().GetUserManager()
	return &userController{
		userManager: um,
	}
}

func (uc *userController) CreateUser(context *gin.Context) {
	var user models.UserRequest

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	u, err := uc.userManager.CreateNewUser(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}

	context.JSON(http.StatusOK, u)
}

func (uc *userController) GetUserById(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse id"})
		return
	}
	user, err := uc.userManager.GetUserById(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": user})
}
