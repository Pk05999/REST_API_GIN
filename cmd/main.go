package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	//Load environmental variable
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
