package config

import (
	"fmt"
	"log/slog"
	"os"
)

type Vars struct {
	JWT_SECRET string
	ENV        string
	DB_URL     string
	REDIS_URL  string
}

func Load() (*Vars, error) {
	// will load all the config variables here
	var variables Vars = Vars{
		JWT_SECRET: getVar("JWT_SECRET", ""),
		ENV:        getVar("ENV", "dev"),
		DB_URL:     getVar("DB_URL", ""),
		REDIS_URL:  getVar("REDIS_URL", ""),
	}

	err := validateVars(variables)

	if err != nil {
		return &Vars{}, err
	}

	return &variables, nil
}

func validateVars(variables Vars) error {
	// Validating jwt
	if variables.JWT_SECRET == "" {
		slog.Error("Jwt secret missing", "JWT_SECRET", variables.JWT_SECRET)
		return fmt.Errorf("Jwt secret is missing")
	}

	if variables.DB_URL == "" {
		slog.Error("Db url missing", "DB_URL", variables.DB_URL)
		return fmt.Errorf("Db url is missing")
	}
	return nil
}

func getVar(varName, defaultVal string) string {
	if value := os.Getenv(varName); value != "" {
		return value
	}
	return defaultVal
}
