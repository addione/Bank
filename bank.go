package main

import (
	"context"
	"fmt"

	"github.com/0x6flab/namegenerator"
	models "github.com/addione/New/models"
)

func loginUser() {

}

func handle() {
	var numberOfUsers int
	fmt.Println("Enter Numebr of users you want to create: ")
	fmt.Scan(&numberOfUsers)
	connectMongo()
	createUsers(numberOfUsers)
	return
}

func createUsers(numberOfUsers int) {
	for i := 0; i < numberOfUsers; i++ {
		createNewUser()
	}
}

func createNewUser() {
	generator := namegenerator.NewGenerator()
	client := getMongoClient("New", "User")
	user := models.User{Name: generator.Generate(), Pass: "pass", Balance: 1000}
	result, err := client.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}

func takeInput() {
	fmt.Println(`
	Welcome to the bank
	what do you want to do?
	1: Check Balance
	2: Deposit Money
	3:Widraw Money
	4: exit `)
}
