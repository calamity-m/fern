package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Debug override
	debug bool

	rootCmd = &cobra.Command{
		Use:   "recorder",
		Short: "Start the recorder server",
		Long:  "tbd",
	}
)

func Execute() {

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Debug Override")
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(configCmd)

	err := rootCmd.Execute()

	if err != nil {
		fmt.Println("Failed to run")
		os.Exit(1)
	}
}
