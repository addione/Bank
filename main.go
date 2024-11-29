package main

func main() {
	handleHttp()

	// server := gin.Default()
	// sdb := src.NewSrcDI().GetBank()

	// server.POST("/new-user", sdb.CreateUser)
	// server.GET("/clean-db", sdb.CleanDb)

	// server.GET("/list-users", sdb.ListUsers)

	// server.Run(":8091")

	// structs()

	// practice.Routines()

	// sd := src.NewSrcDI()

	// var userChoice int
	// fmt.Println("Welcome to the application")
	// fmt.Println(`What defines your role
	//  1. Customer
	//  2. Employee
	//     `)
	// fmt.Scan(&userChoice)
	// fmt.Println(userChoice)
	// switch userChoice {
	// case 2:
	// 	sd.GetBank().Handle()
	// }
}

// func createUser(context *gin.Context) {
// 	manager := manager.NewManagerDIContainer().GetUserManager()
// 	manager.CreateNewUser()
// 	context.JSON()
// }
