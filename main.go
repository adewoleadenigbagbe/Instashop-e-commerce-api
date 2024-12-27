package main

import (
	"log"
	"os"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/cmd"
	"github.com/joho/godotenv"
)

func main() {
	//load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := cmd.NewAPIHost()
	err = host.Start()
	if err != nil {
		os.Exit(1)
	}
}
