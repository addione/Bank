package src

import (
	"fmt"
	"net/http"

	"github.com/addione/New/manager"
	"github.com/gin-gonic/gin"
)

func loginUser() {

}

type bank struct {
	userManager *manager.UserManager
}

func newBank() (b *bank) {
	um := manager.NewManagerDIContainer().GetUserManager()
	return &bank{
		userManager: um,
	}
}

func (b *bank) Handle() {
	var numberOfUsers int
	var userChoice int

	fmt.Println(`What would you like to do
1. Populate Database
2. Clean Database
3. Use Application
		`)

	fmt.Scan(&userChoice)
	switch userChoice {
	case 1:
		fmt.Println("Enter Numebr of users you want to create: ")
		fmt.Scan(&numberOfUsers)
		b.createUsers(numberOfUsers)

	case 2:
		b.userManager.CleanDatabase()
	}

	return
}

func (b *bank) CleanDb(context *gin.Context) {
	b.userManager.CleanDatabase()
	context.JSON(http.StatusOK, "ok")
}

func (b *bank) ListUsers(context *gin.Context) {

	context.JSON(http.StatusOK, b.userManager.ListUsers())
}

func (b *bank) CreateUser(context *gin.Context) {

	context.JSON(http.StatusOK, b.userManager.CreateNewUser())
}

func (b *bank) createUsers(numberOfUsers int) {
	for i := 0; i < numberOfUsers; i++ {
		b.userManager.CreateNewUser()
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
