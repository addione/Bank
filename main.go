package main

import "fmt"

func main() {
	var userChoice int
	fmt.Println("Welcome to the application")
	fmt.Println(`What would you like to do
				1. Populate Database
				2. Use Application
	`)

	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:
		handle()
	}

	handle()
	// const test = "aknaksnk"
	// fmt.Println("this is madsdsn")
	// newfun()
	// fmt.Printf("profit is %.3f \n", calculateProfit(1003, 910, 5.1))
	// fmt.Println(calculateInvestment(1000, 12, 5.5))
}
