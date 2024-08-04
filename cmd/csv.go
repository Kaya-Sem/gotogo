package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
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

func getTodos(path string) ([]todoItem, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", path, err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("Warning: failed to close file %s: %v", path, closeErr)
		}
	}()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV data from %s: %w", path, err)
	}

	todoItems, err := convertCsvToTodos(data)
	if err != nil {
		return nil, fmt.Errorf("failed to convert CSV data to todo items: %w", err)
	}

	return todoItems, nil
}

func getTodos(path string) []todoItem {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var todoItems = convertCsvToTodos(data)
	return todoItems
}

func convertCsvToTodos(data [][]string) []todoItem {
	var todoItems []todoItem
	for i, line := range data {
		if i > 0 { // omit header line
			var item todoItem
			for j, field := range line {
				switch j {
				case 0:
					item.Name = field
				case 1:
					priority, err := strconv.Atoi(field)
					if err != nil {
						log.Fatalf("Invalid priority value: %s", field)
					}
					item.Priority = priority
				case 2:
					isDone, err := strconv.ParseBool(field)
					if err != nil {
						log.Fatalf("Invalid isDone value: %s", field)
					}
					item.isDone = isDone
				case 3:
					item.Timestamp = field
				}
			}
			todoItems = append(todoItems, item)
		}
	}
	return todoItems
}
