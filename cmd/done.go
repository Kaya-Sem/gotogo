/*
Copyright © 2024 Kaya-Sem Van Cauwenberghe <kayasemvc@gmail.com>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done <id>",
	Short: "Mark a todo item as done",
	Long:  `Mark a todo item as done by providing its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID. Please provide a valid integer.")
			return
		}

		err = markItemAsDone(id)
		if err != nil {
			fmt.Printf("Error marking item as done: %v\n", err)
			return
		}

		fmt.Printf("Todo item with ID %d marked as done.\n", id)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func markItemAsDone(id int) error {
	filename, err := getFilePath()
	if err != nil {
		return fmt.Errorf("failed to get file path: %w", err)
	}

	todos, err := readTodos(filename)
	if err != nil {
		return fmt.Errorf("failed to read todos: %w", err)
	}

	found := false
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("todo item with ID %d not found", id)
	}

	// Write updated todos back to the CSV file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	err = writer.Write([]string{"id", "title", "completed", "timestamp"})
	if err != nil {
		return fmt.Errorf("failed to write CSV header: %w", err)
	}

	// Write updated todos
	for _, todo := range todos {
		record := []string{
			strconv.Itoa(todo.ID),
			todo.Title,
			strconv.FormatBool(todo.Completed),
			todo.Timestamp.Format("2006-01-02"),
		}
		err := writer.Write(record)
		if err != nil {
			return fmt.Errorf("failed to write todo item: %w", err)
		}
	}

	return nil
}
