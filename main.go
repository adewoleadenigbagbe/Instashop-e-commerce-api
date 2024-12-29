package main

import (
	"os"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/cmd"
)

func main() {
	host := cmd.NewAPIHost()
	err := host.Start()
	if err != nil {
		os.Exit(1)
	}
}
