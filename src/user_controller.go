package src

import (
	"net/http"

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
