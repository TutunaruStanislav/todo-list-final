package main

import (
	"fmt"

	"gop/pkg/db"
	"gop/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	err = db.Init()
	if err != nil {
		fmt.Printf("DB initialization error: %s", err)
		return
	}

	err = server.Run()
	if err != nil {
		fmt.Printf("Start server error: %s", err.Error())
		return
	}
}
