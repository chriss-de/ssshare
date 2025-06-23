package main

import (
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/chriss-de/ssshare/internal/backend"
	"github.com/chriss-de/ssshare/internal/helpers"
	"github.com/chriss-de/ssshare/internal/server"
)

func main() {
	var err error

	// flags
	flag.Parse()

	// logger
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))

	// backend
	if err = backend.Initialize(); err != nil {
		helpers.FatalError("initializing config", "error", err)
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

//// initializeConfig initialize config with default values and cmd flags and config file
//func initializeConfig() error {
//	viper.SetEnvPrefix("AEBROLES")
//
//	viper.SetDefault("verbose", "2")
//	viper.SetDefault("server.listen_addr", ":8080")
//	viper.SetDefault("server.baseUrl", "/")
//	viper.SetDefault("server.maxHeaderSize", "1mb")
//	viper.SetDefault("server.maxBytesReader", 1048576)
//
//	_ = viper.BindEnv("config", "AEBROLES_CONFIG_FILE")
//
//	viper.SetDefault("database.print_sql", false)
//	viper.SetDefault("database.max_open_conn", 32)
//	viper.SetDefault("database.max_idle_conn", 16)
//	viper.SetDefault("database.max_conn_lifetime", 10*time.Minute)
//
//	viper.SetDefault("app.allow_fake_employee_number", false)
//	viper.SetDefault("app.employee_number_token_field_name", "employee_number")
//
//	viper.BindEnv("database.dsn", "AEBROLES_DATABASE_DSN")
//
//	//viper.SetConfigName(configFileName)
//	viper.AddConfigPath(".")
//	viper.AddConfigPath("./data")
//	viper.AddConfigPath("/config")
//
//	if cfgFileFromEnv := viper.GetString("config"); cfgFileFromEnv != "" {
//		viper.SetConfigFile(cfgFileFromEnv)
//	}
//
//	// Attempt to read the config file, gracefully ignoring errors
//	// caused by a config file not being found. Return an error
//	// if we cannot parse the config file.
//	if err := viper.ReadInConfig(); err != nil {
//		return err
//	}
//
//	return nil
//}
