package models

import "time"

const UserCollectionName = "User"

const userTableName = "users"

const (
	STATUS_NEW      = "new"
	STATUS_ACTIVE   = "active"
	STATUS_INACTIVE = "inactive"
)

type User struct {
	Name        string  `json:"name"`
	ID          int64   `bson:"mysql_id"`
	Email       string  `json:"email"`
	Pass        string  `json:"password"`
	Balance     float64 `json:"balance"`
	PhoneNumber string  `bson:"phone_number"`
	Salutation  string
	Details     Details
	Address     string
	CreatedAt   time.Time `bson:"created_at"`
}

type Details struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
}

type UserMysql struct {
	ID          int64
	Email       string `json:"email"`
	PhoneNumber string
	Password    string `json:"password"`
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserRequest struct {
	Name        string `binding:"required"`
	Email       string `json:"email"`
	PhoneNumber string `bson:"phone_number"`
	Address     string
	Details     Details
	Password    string `binding:"required"`
}

type UserUpdateRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
