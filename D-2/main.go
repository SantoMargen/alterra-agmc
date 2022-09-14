package main

import (
	"D-2/routes"
)

func main() {
	// create a new echo instance
	app := routes.New()

	// Route / to handler function

	// password := []byte("Hello, password")

	// // Hashing the password with the default cost of 10
	// hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(hashedPassword))

	// // Comparing the password with the hash
	// err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	// fmt.Println(err) // nil means it is a match

	app.Start(":8080")
}
