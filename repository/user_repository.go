package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/0x6flab/namegenerator"
	"github.com/addione/New/dependencies"
	"github.com/addione/New/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	mongoClient *mongo.Collection
}

func NewUserRepository() *UserRepo {
	cm := dependencies.NewCommonMongo().GetMongoClient("New", "User")
	return &UserRepo{
		mongoClient: cm,
	}
}

func CreateUser() {
	return
}

func (u *UserRepo) CreateNewUser() {
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
