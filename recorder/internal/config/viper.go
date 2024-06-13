package config

import (
	"errors"
	"log/slog"
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	ConfigFileSource string
	Env              string `mapstructure:"env" json:"environment,omitempty"`
	Server           server
	Tls              tls
	Logging          logging
}

type server struct {
	Host string `mapstructure:"host" json:"host,omitempty"`
	Port int    `mapstructure:"port" json:"port,omitempty"`
}

type tls struct {
	Enabled        bool   `mapstructure:"enabled" json:"enabled,omitempty"`
	Certificate    string `mapstructure:"certificate" json:"certificate,omitempty"`
	CertificateKey string `mapstructure:"certificate_key" json:"certificate_key,omitempty"`
}

type logging struct {
	Level      slog.Level `mapstructure:"level" json:"level,omitempty"`
	Structured bool       `mapstructure:"structured" json:"structured,omitempty"`
	AddSource  bool       `mapstructure:"add_source" json:"add_source,omitempty"`
}

func newBaseConfig() Config {
	return Config{
		Env: "dev",
		Server: server{
			Host: "localhost",
			Port: 5005,
		},
		Tls: tls{
			Enabled: false,
		},
		Logging: logging{
			Level:      slog.LevelInfo,
			Structured: true,
			AddSource:  true,
		},
	}
}

// Returns a new configuration object built from a sane default, overriden by a sourced
// viper yaml file on disk.
//
// cfgPath is the full path to the configuration file to use. If empty or not supplied, viper will look for
// and create a configuration file instead.
func NewConfig(cfgPath string) (Config, error) {
	cfg := newBaseConfig()

	vip := viper.New()

	if cfgPath == "" {
		// Define a place to find and source a config file for recorder

		vip.SetConfigName("recorder")
		vip.SetConfigType("yaml")

		// Ensure the user has a configuration directiory
		cfgDir, err := os.UserConfigDir()
		if err != nil {
			return Config{}, err
		}

		// Append a grouping to the configuration directory in case we have many services under this monorepo
		cfgDir = cfgDir + "/fern"

		// Ensure the directory exists if we haven't been able to find a valid directory
		if err := os.MkdirAll(cfgDir, 0766); err != nil {
			return Config{}, err
		}

		vip.AddConfigPath(cfgDir)
	} else {
		vip.SetConfigFile(cfgPath)
	}

	// Environment variable override
	vip.AutomaticEnv()
	vip.SetEnvPrefix("FERN_RECORDER")
	vip.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// Set values to check for
	vip.SetDefault("env", cfg.Env)
	vip.SetDefault("server.host", cfg.Server.Host)
	vip.SetDefault("server.port", cfg.Server.Port)
	vip.SetDefault("tls.enabled", cfg.Tls.Enabled)
	vip.SetDefault("tls.certificate", cfg.Tls.Certificate)
	vip.SetDefault("tls.certificate_key", cfg.Tls.CertificateKey)
	vip.SetDefault("logging.level", cfg.Logging.Level)
	vip.SetDefault("logging.structured", cfg.Logging.Structured)
	vip.SetDefault("logging.add_source", cfg.Logging.AddSource)

	// Parse in file

	// Try to read in the file, and create it if it wasn't found
	if err := vip.ReadInConfig(); err != nil {
		var actual viper.ConfigFileNotFoundError
		if errors.As(err, &actual) {
			if err := vip.SafeWriteConfig(); err != nil {
				return Config{}, err
			}
		} else {
			return Config{}, err
		}
	}

	if err := vip.Unmarshal(&cfg, viper.DecodeHook(mapstructure.TextUnmarshallerHookFunc())); err != nil {
		return Config{}, err
	}

	cfg.ConfigFileSource = vip.ConfigFileUsed()

	return cfg, nil
}
