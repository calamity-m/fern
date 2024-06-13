package cmd

import (
	"fmt"
	"os"

	"github.com/calamity-m/fern/recorder/internal/config"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start serving",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Serve")

			cfg, err := config.NewConfig("")

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(cfg)
		},
	}
)
