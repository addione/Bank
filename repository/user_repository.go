package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/addione/New/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	userCollection *mongo.Collection
	userTable      *sql.DB
}

func newUserRepository(rdi *RepositoryDIContainer) *UserRepo {
	return &UserRepo{
		userCollection: rdi.DependenciesDI.GetMongoCollection(DBName, models.UserCollectionName),
		userTable:      rdi.DependenciesDI.GetMysql(MySQLDBName),
	}
}

func (u *UserRepo) CreateNewUser(user *models.User) {
	userId, _ := u.insertIntoMysqlTable(user)
	user.ID = userId
	fmt.Println(user)
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

func (u *UserRepo) insertIntoMysqlTable(user *models.User) (int64, error) {
	query := "INSERT into users(email, phone_number, password, status) VALUES(?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := u.userTable.PrepareContext(ctx, query)

	res, err := stmt.ExecContext(ctx, user.Email, user.Balance, user.Pass, models.STATUS_NEW)
	fmt.Println(err, "............")

	return res.LastInsertId()
}
