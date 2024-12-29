package cmd

import (
	"os"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/server"
	"github.com/spf13/cobra"
)

func serveApiCommand() *cobra.Command {
	var apiCmd = &cobra.Command{
		Use:   "serveapi",
		Short: "Serve the API on the Specified host",
		Run: func(cmd *cobra.Command, args []string) {
			//set the application env
			os.Setenv("APP_ENV", environment)
			server.InitializeAPI()
		},
	}

	return apiCmd
}
