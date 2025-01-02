package cmd

import (
	"github.com/spf13/cobra"
)

var (
	environment string
)

type ApiHost struct {
	rootCmd *cobra.Command
}

func NewAPIHost() *ApiHost {
	tc := &ApiHost{
		rootCmd: &cobra.Command{
			Use:   "instashop",
			Short: "ApiHost CLI",
			// no need to provide the default cobra completion command
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
		},
	}

	return tc
}

func (host *ApiHost) Start() error {
	s := serveApiCommand()
	host.rootCmd.AddCommand(s)
	s.Flags().StringVarP(&environment, "APP_ENV", "a", "", "Set the environment (e.g.development,staging, production)")

	return host.execute()
}

func (host *ApiHost) execute() error {
	err := host.rootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
