package main

import "fmt"

func main() {

	structs()

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
		bank := newBank()
		bank.Handle()
	}

	// const test = "aknaksnk"
	// fmt.Println("this is madsdsn")
	// newfun()
	// fmt.Printf("profit is %.3f \n", calculateProfit(1003, 910, 5.1))
	// fmt.Println(calculateInvestment(1000, 12, 5.5))
}
