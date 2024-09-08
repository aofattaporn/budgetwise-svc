package cmd

import (
	"fmt"

	"github.com/goproject/configs"
	"github.com/goproject/server"
	"github.com/spf13/cobra"
)

// command startUpCmd usage
var startUpCmd = &cobra.Command{
	Use:   "startup",
	Short: "Startup Server",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Loading config
		config, err := configs.LoadConfig(configFile)
		if err != nil {
			return fmt.Errorf("load config error: %+v", err)
		}

		// initial server
		s, err := server.NewFiberServer(config)
		if err != nil {
			return err
		}

		// start mmy servers
		s.Start()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(startUpCmd)
}
