package cmd

import (
	"fmt"
	"os"

	"github.com/calamity-m/fern/recorder/internal/config"
	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Display the configuration that will be evaluated by the recorder server when running. Can also be used to generate config before first run if wanted.",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.NewConfig("")

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			text := `Config file used: %s

Configuration successfully parsed:

GENERAL:
    Environment: %s

SERVER:
    Server Host - %s
    Server Port - %d

TLS:
    TLS enabled         - %t
    TLS certificate     - %s
    TLS certificate key - %s

LOGGING:
    Logging level      - %s
    Logging structured - %t
    Logging add source - %t
`

			fmt.Printf(text,
				cfg.ConfigFileSource,
				cfg.Env,
				cfg.Server.Host,
				cfg.Server.Port,
				cfg.Tls.Enabled,
				cfg.Tls.Certificate,
				cfg.Tls.CertificateKey,
				cfg.Logging.Level,
				cfg.Logging.Structured,
				cfg.Logging.AddSource)
		},
	}
)
