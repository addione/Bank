package manager

import (
	"time"

	"github.com/0x6flab/namegenerator"
	"github.com/addione/New/models"
	"github.com/addione/New/repository"
)

type UserManager struct {
	userRepo *repository.UserRepo
}

func newUserManager(mdi *managerDIContainer) *UserManager {
	return &UserManager{
		userRepo: mdi.repositoryDIContainer.GetUserRepo(),
	}
}

func (um *UserManager) CreateNewUser() {

	generator := namegenerator.NewGenerator()
	name := generator.Generate()
	user := models.User{
		Name:  name,
		Email: name + `@gmail.com`,
		Pass:  "pass", Balance: 1000,
		Details: models.Details{
			FirstName: name,
			LastName:  name,
		},
		CreatedAt: time.Now(),
	}
	um.userRepo.CreateNewUser(&user)
}

func (um *UserManager) CleanDatabase() {
	um.userRepo.CleanDatabase()
}
