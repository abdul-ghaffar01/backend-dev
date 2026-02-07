package main

import (
	"fmt"
	"log/slog"
	"os"
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

}
