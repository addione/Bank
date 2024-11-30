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

func (u *UserRepo) CreateNewUser(user *models.User) *models.User {
	userId, _ := u.insertIntoMysqlTable(user)
	user.ID = userId
	fmt.Println(user)
	_, err := u.userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}

	return user
}

func (u *UserRepo) CleanDatabase() {
	filter := bson.D{}
	result, _ := u.userCollection.DeleteMany(context.TODO(), filter)
	u.cleanMysqlDB()
	fmt.Println(result)
}

func (u *UserRepo) GetAllUsers() []*models.User {
	var users []*models.User
	filter := bson.D{}
	result, err := u.userCollection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	err = result.All(context.TODO(), &users)
	if err != nil {
		panic(err)
	}

	return users
}

func (u *UserRepo) UpdateUserByID(userId int64, ur *models.UserUpdateRequest) error {
	query := "UPDATE users set email =?, phone_number =? WHERE id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelfunc()
	stmt, err := u.userTable.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	_, err = stmt.ExecContext(ctx, ur.Email, ur.PhoneNumber, userId)
	stmt.Close()
	return err
}

func (u *UserRepo) GetUserByEmail(email string) (*models.UserMysql, error) {
	row := u.userTable.QueryRow(`SELECT id, created_at from users WHERE email = ?`, email)
	var user models.UserMysql
	err := row.Scan(&user.ID, &user.CreatedAt)
	return &user, err
}

func (u *UserRepo) GetUserById(id int64) (*models.UserMysql, error) {
	row := u.userTable.QueryRow(`SELECT id,email, created_at from users WHERE id = ?`, id)
	var user models.UserMysql
	err := row.Scan(&user.ID, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) insertIntoMysqlTable(user *models.User) (int64, error) {
	query := "INSERT into users(email, phone_number, password, status) VALUES(?, ?, ?, ?) ON DUPLICATE KEY UPDATE phone_number=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelfunc()

	stmt, err := u.userTable.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	res, _ := stmt.ExecContext(ctx, user.Email, user.PhoneNumber, user.Pass, models.STATUS_NEW, user.PhoneNumber)
	stmt.Close()
	return res.LastInsertId()
}

func (u *UserRepo) cleanMysqlDB() {
	query := "DELETE from users "
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, _ := u.userTable.PrepareContext(ctx, query)
	stmt.ExecContext(ctx)
}
