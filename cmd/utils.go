package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// Purges completed items older than a day old
func purgeCompletedItems() error {
	filename, err := getFilePath()
	if err != nil {
		return fmt.Errorf("failed to get file path: %w", err)
	}

	todos, err := readTodos(filename)
	if err != nil {
		return fmt.Errorf("failed to read todos: %w", err)
	}

	var updatedTodos []todoItem
	now := time.Now()

	for _, todo := range todos {
		// Keep the item if it's not completed or if it's less than a day old
		if !todo.Completed || now.Sub(todo.Timestamp) < 24*time.Hour {
			updatedTodos = append(updatedTodos, todo)
		}
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
	for _, todo := range updatedTodos {
		record := []string{
			fmt.Sprintf("%d", todo.ID),
			todo.Title,
			fmt.Sprintf("%t", todo.Completed),
			todo.Timestamp.Format("2006-01-02"),
		}
		err := writer.Write(record)
		if err != nil {
			return fmt.Errorf("failed to write todo item: %w", err)
		}
	}

	return nil
}

func formatRelativeDate(t time.Time) string {
	now := time.Now()
	today := now.Truncate(24 * time.Hour)
	givenDay := t.Truncate(24 * time.Hour)

	if givenDay.Equal(today) {
		return "Today"
	}

	if givenDay.Equal(today.Add(-24 * time.Hour)) {
		return "Yesterday"
	}

	days := int(today.Sub(givenDay).Hours() / 24)
	return fmt.Sprintf("%d days ago", days)
}
