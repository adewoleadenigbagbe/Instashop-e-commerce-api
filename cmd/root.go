package cmd

import (
	"github.com/spf13/cobra"
)

type ApiHost struct {
	rootCmd *cobra.Command
}

func NewAPIHost() *ApiHost {
	tc := &ApiHost{
		rootCmd: &cobra.Command{
			Use:   "serveapi",
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
	host.rootCmd.AddCommand(serveApiCommand())

	return host.execute()
}

func (host *ApiHost) execute() error {
	err := host.rootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
