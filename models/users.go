package models

import "time"

const UserCollectionName = "User"

type User struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Pass      string  `json:"password"`
	Balance   float64 `json:"balance"`
	Details   Details
	CreatedAt time.Time `bson:"created_at"`
}

type Details struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
}
