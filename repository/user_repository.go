package repository

import (
	"context"
	"fmt"

	"github.com/addione/New/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	userCollection *mongo.Collection
}

func newUserRepository(rdi *RepositoryDIContainer) *UserRepo {
	return &UserRepo{
		userCollection: rdi.DependenciesDI.GetMongoCollection(DBName, models.UserCollectionName),
	}
}

func (u *UserRepo) CreateNewUser(user *models.User) {
	result, err := u.userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}

func (u *UserRepo) CleanDatabase() {
	filter := bson.D{}
	result, _ := u.userCollection.DeleteMany(context.TODO(), filter)
	fmt.Println(result)
}
