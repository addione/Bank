package manager

import "github.com/addione/New/repository"

type managerDIContainer struct {
	repositoryDIContainer *repository.RepositoryDIContainer
	userManager           *UserManager
}

func NewManagerDIContainer() *managerDIContainer {
	mdi := &managerDIContainer{}
	mdi.repositoryDIContainer = repository.NewRepositoryDiContainer()
	mdi.userManager = newUserManager(mdi)
	return mdi
}

func (mdi *managerDIContainer) GetUserManager() *UserManager {
	return mdi.userManager
}
