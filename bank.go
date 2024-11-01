package main

import (
	"context"
	"fmt"

	"github.com/0x6flab/namegenerator"
)

func loginUser() {

}

func handle() {
	createNewUser()
}

func createNewUser() {
	generator := namegenerator.NewGenerator()
	client := getMongoClient("New", "User")
	user := User{Name: generator.Generate(), Pass: "pass", Balance: 1000}
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
