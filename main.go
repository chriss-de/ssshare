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
	viper.SetDefault("server.base_url", "/")
	viper.SetDefault("server.max_header_size", "1mb")
	viper.SetDefault("server.max_bytes_reader", 1048576)

	viper.SetDefault("server.cors.allowed_origins", []string{"*"})
	viper.SetDefault("server.cors.allowed_methods", []string{"*"})
	viper.SetDefault("server.cors.allowed_headers", []string{"*"})
	viper.SetDefault("server.cors.exposed_headers", []string{"*"})
	viper.SetDefault("server.cors.max_age", 300)
	viper.SetDefault("server.cors.allow_credentials", true)
	viper.SetDefault("server.cors.debug", false)

	viper.SetDefault("shares.url_path_prefix", "/s")
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
