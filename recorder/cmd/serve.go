package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/calamity-m/fern/recorder/internal/config"
	"github.com/calamity-m/fern/recorder/pkg/logging"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start serving",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.NewConfig("")

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			logger := slog.New(logging.New(
				logging.WithEnvironment(cfg.Env),
				logging.WithBaseHandler(cfg.Logging.Structured,
					cfg.Logging.Level,
					cfg.Logging.AddSource,
				)))

			logger.Info("Initialized logging and configuration")
		},
	}
)
