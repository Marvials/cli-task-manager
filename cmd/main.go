package main

import (
	"log"

	"github.com/Marvials/cli-task-manager/cmd/root"
	_ "github.com/Marvials/cli-task-manager/cmd/table"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load the .env file: ", err)
	}

	root.Execute()
}
