package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "abarobotics-test/docs"
	"abarobotics-test/src/api"
	"abarobotics-test/src/kernel"
	"abarobotics-test/toolkit/config"
	"abarobotics-test/toolkit/db"
	"abarobotics-test/toolkit/logger"
)

// @title Abarobotics API
// @version 1.0
// @description API technical test Abarobotics
// @host localhost:8000
// @BasePath /
// @schemes http
func main() {
	var err error

	// load .env file
	if os.Getenv("APP_ENV") == "" {
		err = godotenv.Load(".env")
		if err != nil {
			log.Fatalf("ERROR load env file : %s", err.Error())
		}
	}

	ctx, cancel := config.NewRuntimeContext()
	defer func() {
		cancel()

		if err != nil {
			log.Printf("found error : %s", err.Error())
		}
	}()

	setDefaultTimezone()

	// setup logger
	logger.NewLogger()

	// setup database
	dbx, db, err := db.NewDatabase()
	if err != nil {
		log.Printf("ERROR setup database : %s", err.Error())
		return
	}

	// setup module
	k := kernel.NewKernel(dbx, db)

	// run fiber http
	api.RunFiberServer(ctx, k)
}

func setDefaultTimezone() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		loc = time.Now().Location()
	}

	time.Local = loc
}
