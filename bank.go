package main

import (
	"fmt"

	"github.com/addione/New/repository"
)

func loginUser() {

}

type bank struct {
	userRepository *repository.UserRepo
}

func newBank() (b *bank) {
	ur := repository.NewUserRepository()

	return &bank{
		userRepository: ur,
	}
}

func (b *bank) Handle() {
	var numberOfUsers int
	var userChoice int

	fmt.Println(`What would you like to do
	1. Populate Database
	2. Clean Database
	2. Use Application
		`)

	fmt.Scan(&userChoice)
	switch userChoice {
	case 1:
		fmt.Println("Enter Numebr of users you want to create: ")
		fmt.Scan(&numberOfUsers)
		b.createUsers(numberOfUsers)

	case 2:
		b.userRepository.CleanDatabase()
	}

	return
}

func (b *bank) createUsers(numberOfUsers int) {
	for i := 0; i < numberOfUsers; i++ {
		b.userRepository.CreateNewUser()
	}
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
