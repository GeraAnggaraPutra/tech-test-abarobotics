package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"abarobotics-test/src/api"
	"abarobotics-test/src/kernel"
	"abarobotics-test/toolkit/config"
	"abarobotics-test/toolkit/db"
	"abarobotics-test/toolkit/logger"
)

// @title           ABAROBOTICS API
// @version         1.0
// @description     Technical Test
// @contact.name    Tim Developer
// @contact.email   anggaragera@gamil.com
// @host            localhost:8000
// @BasePath        /
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
