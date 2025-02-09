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
	Long:  ``,
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

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0600)
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
	var lastID int = 0

	if len(records) > 0 {
		lastID = getLastId(records)
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

func getLastId(records [][]string) int {
	lastRecord := records[len(records)-1]
	lastID, _ := strconv.Atoi(lastRecord[0])

	return lastID
}
