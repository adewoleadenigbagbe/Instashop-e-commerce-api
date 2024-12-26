package main

import (
	"log"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/server"
	"github.com/joho/godotenv"
)

func main() {
	//load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server.InitializeAPI()
}
