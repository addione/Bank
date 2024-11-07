package main

import (
	"fmt"

	"github.com/addione/New/src"
)

func main() {

	// structs()

	sd := src.NewSrcDI()

	var userChoice int
	fmt.Println("Welcome to the application")
	fmt.Println(`What defines your role
1. Customer
2. Employee
	`)
	fmt.Scan(&userChoice)
	fmt.Println(userChoice)
	switch userChoice {
	case 2:
		sd.GetBank().Handle()
	}
}
