package manager

import (
	"github.com/addione/New/helpers"
	"github.com/addione/New/repository"
)

type ManagerDIContainer struct {
	repositoryDIContainer *repository.RepositoryDIContainer
	userManager           *UserManager
	helpers               *helpers.HelpersDIContainer
}

func NewManagerDIContainer() *ManagerDIContainer {
	mdi := &ManagerDIContainer{
		repositoryDIContainer: repository.NewDIContainer(),
		helpers:               helpers.NewHelpersDIContainer(),
	}
	mdi.userManager = newUserManager(mdi)
	return mdi
}

func (mdi *ManagerDIContainer) GetUserManager() *UserManager {
	return mdi.userManager
}
