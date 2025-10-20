package main

import (
	"fmt"
	"log"

	"github.com/Marvials/cli-task-manager/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load the .env file: ", err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	fmt.Println("Connection was established successfully: ", db)

}
