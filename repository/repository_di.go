package repository

import (
	"github.com/addione/New/dependencies"
)

const DBName = "New"

type RepositoryDIContainer struct {
	userRepository *UserRepo
	mongoClient    *dependencies.CommonMongo
}

func NewRepositoryDiContainer() *RepositoryDIContainer {
	cm := dependencies.NewCommonMongo()

	rdi := &RepositoryDIContainer{
		mongoClient: cm,
	}
	rdi.userRepository = newUserRepository(rdi)
	return rdi
}

func (di *RepositoryDIContainer) GetUserRepo() *UserRepo {
	return di.userRepository
}
