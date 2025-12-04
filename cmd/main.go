package main

import (
	"log"
	"os"
	"path/filepath"

	_ "github.com/Marvials/cli-task-manager/cmd/add"
	_ "github.com/Marvials/cli-task-manager/cmd/change"
	_ "github.com/Marvials/cli-task-manager/cmd/delete"
	_ "github.com/Marvials/cli-task-manager/cmd/get"
	_ "github.com/Marvials/cli-task-manager/cmd/list"
	"github.com/Marvials/cli-task-manager/cmd/root"
	_ "github.com/Marvials/cli-task-manager/cmd/table"

	"github.com/joho/godotenv"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not find user directory: ", err)
	}

	envPath := filepath.Join(homeDir, ".config", ".task-manager.env")

	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading the .env file ", err)
	}

	root.Execute()
}
