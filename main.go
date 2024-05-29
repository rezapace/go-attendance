package main

import (
	"log"
	"os"

	bootsrapper "presensee_project/bootstrapper"
	"presensee_project/config"
	"presensee_project/middleware"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
)

func init() {
	if os.Getenv("ENV") == "production" {
		return
	}

	//	load env variables from .env file for local development
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	env := config.LoadConfig()

	db := config.InitDB()
	e := echo.New()
	middleware.Logmiddleware(e)

	bootsrapper.InitController(e, db, env)

	e.Logger.Fatal(e.Start(":80"))

}
