package main

import (
	"D-2/config"
	m "D-2/middlewares"
	"D-2/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	config.InitDB()

	// create a new echo instance
	app := routes.New()
	m.LogMiddleware(app)
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
	app.Logger.Fatal(app.Start(":8080"))
}
