package repository

import (
	"github.com/addione/New/dependencies"
)

const DBName = "New"

type repositorydiContainer struct {
	userRepository *UserRepo
	mongoClient    *dependencies.CommonMongo
}

func NewRepositoryDiContainer() *repositorydiContainer {
	cm := dependencies.NewCommonMongo()

	rdi := &repositorydiContainer{
		mongoClient: cm,
	}
	rdi.userRepository = newUserRepository(rdi)
	return rdi
}

func (di *repositorydiContainer) GetUserRepo() *UserRepo {
	return di.userRepository
}
