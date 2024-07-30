package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/go-project-template/internal/adapter/config"
	"github.com/go-project-template/internal/adapter/logger"
)

func main() {
	fmt.Println("New Project")
	// Load environment variables
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	// Set logger
	logger.Set(config.App)

	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)
}
