package main

import (
	"log"

	"gop/pkg/db"
	"gop/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	// load the .env file with environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error while loading .env: %v", err)
	}

	// initialize the SQlite file DB
	db, err := db.Init()
	if err != nil {
		log.Fatalf("DB initialization error: %s", err)
	}

	// run the server
	err = server.Run(db)
	if err != nil {
		log.Fatalf("Start server error: %s", err.Error())
	}
}
