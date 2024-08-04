/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an item to your todo-list",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a todo item.")
			return
		}

		todoString := args[0]
		err := createTodo(todoString)
		if err != nil {
			fmt.Printf("Error adding todo item: %v\n", err)
			return
		}

		fmt.Println("Todo item added successfully.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func createTodo(todo string) error {
	filePath, err := getFilePath()
	if err != nil {
		return fmt.Errorf("could not get file path: %v", err)
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	// Create a new reader to read existing records
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("could not read CSV file: %v", err)
	}

	// Get the last ID
	var lastID int
	if len(records) > 0 {
		lastRecord := records[len(records)-1]
		lastID, err = strconv.Atoi(lastRecord[0])
		if err != nil {
			return fmt.Errorf("could not parse last ID: %v", err)
		}
	}

	// Increment the ID for the new record
	newID := lastID + 1

	// Get the current date in the desired format
	currentDate := time.Now().Format("2006-01-02")

	// Create the new record
	newRecord := []string{
		strconv.Itoa(newID),
		todo,
		"false",
		currentDate,
	}

	// Create a writer to append the new record
	writer := csv.NewWriter(file)
	err = writer.Write(newRecord)
	if err != nil {
		return fmt.Errorf("could not write to CSV file: %v", err)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("could not flush writer: %v", err)
	}

	return nil
}
