package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	configFile string
	rootCmd    = &cobra.Command{
		Use:   "goproject-ci-batch",
		Short: "goproject-ci-batch",
	}
)

func init() {

	// add flags into root command
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "env/config.yaml", "config file (default is env/config.yaml)")
}

func Excecute() {

	// root command is when excute program
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
