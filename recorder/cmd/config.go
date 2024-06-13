package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Display the configuration that will be evaluated by the recorder server when running",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Muh config")

			if debug {
				fmt.Println("Debuggin")
			}
		},
	}
)
