package repository

import (
	"github.com/addione/New/dependencies"
)

const DBName = "bank"
const MySQLDBName = "bank"

type RepositoryDIContainer struct {
	userRepository *UserRepo
	DependenciesDI *dependencies.DependenciesDI
}

func NewDIContainer() *RepositoryDIContainer {
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
