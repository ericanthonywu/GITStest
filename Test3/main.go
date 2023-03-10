package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"laundry/Lib"
	"laundry/Middleware"
	"laundry/Model/Database"
	"laundry/Route"
	"os"
)

func main() {
	e := echo.New()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	Lib.InitAll()

	// run migration
	err := Lib.DB.AutoMigrate(
		&Database.Book{},
		&Database.Publisher{},
		&Database.Author{},
		&Database.User{},
	)
	if err != nil {
		panic(err)
	}

	Middleware.Init(e)
	Route.Init(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
