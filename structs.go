package main

import (
	"fmt"
	"time"

	"github.com/addione/New/models"
)

type User struct {
	Name      string
	LastName  string
	BirthDate time.Time
}

func structs() {
	user := User{
		Name:      "Akash",
		LastName:  "Tyagi",
		BirthDate: time.Now(),
	}
	user1 := models.User{
		Name:    "test",
		Pass:    "BAkjas",
		Balance: 10.10,
	}
	s := "ansakn"
	user.incrementData()

	upateData(user, &user1, &s)
	fmt.Println(user1, s)
}

func upateData(u User, u1 *models.User, s *string) {
	u.BirthDate = time.Date(1993, time.April, 12, 12, 42, 0, 0, time.UTC)
	u1.Balance = 19212
	fmt.Println(u, u1, *s)
	*s = "asnaknsssss"
}

func (u *User) incrementData() {
	u.Name += "ansjans"
	u.LastName += "asnasnajnsjansjansjansjan"
}
