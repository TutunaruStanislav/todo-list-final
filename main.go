package main

import (
	"fmt"

	"gop/pkg/db"
	"gop/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	// load the .env file with environment variables
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	// initialize the SQlite file DB
	db, err := db.Init()
	if err != nil {
		fmt.Printf("DB initialization error: %s", err)
		return
	}

	// run the server
	err = server.Run(db)
	if err != nil {
		fmt.Printf("Start server error: %s", err.Error())
		return
	}
}
