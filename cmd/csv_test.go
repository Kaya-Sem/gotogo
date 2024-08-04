package cmd

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

func TestGetFileFromConfigHome(t *testing.T) {
	// Save the original env var to restore it later
	originalEnv := os.Getenv(EnvConfigHome)
	defer os.Setenv(EnvConfigHome, originalEnv)

	// Get current user for home directory path
	currentUser, err := user.Current()
	if err != nil {
		t.Fatalf("Failed to get current user: %v", err)
	}

	tests := []struct {
		name           string
		envValue       string
		expectedSubdir string
	}{
		{
			name:           "With XDG_CONFIG_HOME set",
			envValue:       "/custom/config/path",
			expectedSubdir: filepath.Join("/custom/config/path", AppFolder),
		},
		{
			name:           "Without XDG_CONFIG_HOME set",
			envValue:       "",
			expectedSubdir: filepath.Join(currentUser.HomeDir, DefaultConfigFolder, AppFolder),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(EnvConfigHome, tt.envValue)

			got, err := getFilePath()
			if err != nil {
				t.Fatalf("getFileFromConfigHome() error = %v", err)
			}

			if !filepath.IsAbs(got) {
				t.Errorf("getFileFromConfigHome() returned non-absolute path: %v", got)
			}

			if got != filepath.Join(tt.expectedSubdir, TodoFileName) {
				t.Errorf("getFileFromConfigHome() = %v, want %v", got, filepath.Join(tt.expectedSubdir, TodoFileName))
			}
		})
	}
}

func TestGetFileFromConfigHomeError(t *testing.T) {
	// This test remains the same as before
	_, err := getFilePath()
	if err != nil {
		t.Fatalf("getFileFromConfigHome() unexpected error: %v", err)
	}
}
