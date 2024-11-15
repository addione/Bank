package manager

import "github.com/addione/New/repository"

type ManagerDIContainer struct {
	repositoryDIContainer *repository.RepositoryDIContainer
	userManager           *UserManager
}

func NewManagerDIContainer() *ManagerDIContainer {
	mdi := &ManagerDIContainer{}
	mdi.repositoryDIContainer = repository.NewDIContainer()
	mdi.userManager = newUserManager(mdi)
	return mdi
}

func (mdi *ManagerDIContainer) GetUserManager() *UserManager {
	return mdi.userManager
}
