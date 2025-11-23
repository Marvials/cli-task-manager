package main

import (
	"log"

	_ "github.com/Marvials/cli-task-manager/cmd/add"
	_ "github.com/Marvials/cli-task-manager/cmd/change"
	_ "github.com/Marvials/cli-task-manager/cmd/get"
	_ "github.com/Marvials/cli-task-manager/cmd/list"
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
