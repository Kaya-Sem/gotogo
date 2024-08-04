package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

const (
	// XDG environment variable name
	EnvConfigHome = "XDG_CONFIG_HOME"

	// Default configuration paths
	DefaultConfigFolder = ".config"
	AppFolder           = "gotogo"
	TodoFileName        = "todo.csv"
)

func getFileFromConfigHome() (string, error) {
	configHome := os.Getenv(EnvConfigHome)
	if configHome == "" {
		usr, err := user.Current()
		if err != nil {
			return "", fmt.Errorf("failed to get current user: %w", err)
		}
		configHome = filepath.Join(usr.HomeDir, DefaultConfigFolder)
	}
	csvPath := filepath.Join(configHome, AppFolder, TodoFileName)
	return csvPath, nil
}
