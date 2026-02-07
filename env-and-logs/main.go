package main

import (
	"env-and-logs/config"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Started structured logging")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	slog.Info("Started server", "port", 8080)

	logger = slog.With(
		"Service", "auth",
		"version", "1.0.0")
	logger.Info("login successfull", "user_id", 13423)
	logger.Error("Token not found")

	fmt.Println("Environment variables and config management start here-------------------------")

	// load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env")
	}
	
	cfg, err := config.Load()

	if err != nil {
		slog.Error("Config not loaded properly", "error", err)
	}

	fmt.Println(cfg)
}
