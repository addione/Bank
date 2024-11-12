package repository

import (
	"github.com/addione/New/dependencies"
)

const DBName = "New"

type RepositoryDIContainer struct {
	userRepository *UserRepo
	DependenciesDI *dependencies.DependenciesDI
}

func NewRepositoryDiContainer() *RepositoryDIContainer {
	ddi := dependencies.NewDependenciesDIProvider()

	rdi := &RepositoryDIContainer{
		DependenciesDI: ddi,
	}
	rdi.userRepository = newUserRepository(rdi)
	return rdi
}

func (di *RepositoryDIContainer) GetUserRepo() *UserRepo {
	return di.userRepository
}
