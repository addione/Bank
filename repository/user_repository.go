package repository

import (
	"context"
	"fmt"

	"github.com/addione/New/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	mongoClient *mongo.Collection
}

func newUserRepository(rdi *RepositoryDIContainer) *UserRepo {

	return &UserRepo{
		mongoClient: rdi.mongoClient.GetMongoClient(DBName, models.UserCollectionName),
	}
}

func (u *UserRepo) CreateNewUser(user *models.User) {
	result, err := u.mongoClient.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}

func (u *UserRepo) CleanDatabase() {
	filter := bson.D{}
	result, _ := u.mongoClient.DeleteMany(context.TODO(), filter)
	fmt.Println(result)
}
