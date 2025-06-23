package main

import (
	"errors"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/chriss-de/ssshare/internal/backend"
	"github.com/chriss-de/ssshare/internal/helpers"
	"github.com/chriss-de/ssshare/internal/server"

	"github.com/spf13/viper"
)

func main() {
	var err error

	// flags
	flag.Parse()

	if err = initializeConfig(); err != nil {
		helpers.FatalError("initializing config", "error", err)
	}

	// logger
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))

	// backend
	if err = backend.Initialize(); err != nil {
		helpers.FatalError("initializing backend", "error", err)
	}

	// start server
	if err = server.Initialize(); err != nil {
		helpers.FatalError("error initialize server", "error", err.Error())
	}

	// Wait for a SIGINT or SIGTERM signal to gracefully shut down the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}

// initializeConfig initialize config with default values and cmd flags and config file
func initializeConfig() error {
	viper.SetEnvPrefix("SSSHARE")

	viper.SetDefault("verbose", "2")
	viper.SetDefault("server.listen_addr", ":8080")
	viper.SetDefault("server.baseUrl", "/")
	viper.SetDefault("server.maxHeaderSize", "1mb")
	viper.SetDefault("server.maxBytesReader", 1048576)

	viper.SetDefault("shares.backend", "file")
	viper.SetDefault("shares_backend.file.file", "shares.yaml")

	_ = viper.BindEnv("config", "CONFIG_FILE")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./data")
	viper.AddConfigPath("/config")

	if cfgFileFromEnv := viper.GetString("config"); cfgFileFromEnv != "" {
		viper.SetConfigFile(cfgFileFromEnv)
	}

	// Attempt to read the config file, gracefully ignoring errors
	// caused by a config file not being found. Return an error
	// if we cannot parse the config file.
	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return err
		}
	}

	return nil
}
