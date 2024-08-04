package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

var FieldIndexes = map[string]int{
	"id":        0,
	"title":     1,
	"completed": 2,
	"timestamp": 3,
}

type todoItem struct {
	ID        int
	Title     string
	Completed bool
	Timestamp time.Time
}

func readTodos(filename string) ([]todoItem, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	var todos []todoItem
	for i, record := range records {
		if i == 0 { // Skip header
			continue
		}

		id, err := strconv.Atoi(record[FieldIndexes["id"]])
		if err != nil {
			return nil, fmt.Errorf("invalid ID at row %d: %w", i+1, err)
		}

		completed, err := strconv.ParseBool(record[FieldIndexes["completed"]])
		if err != nil {
			return nil, fmt.Errorf("invalid completed status at row %d: %w", i+1, err)
		}

		timestamp, err := time.Parse("2006-01-02", record[FieldIndexes["timestamp"]])
		if err != nil {
			return nil, fmt.Errorf("invalid timestamp at row %d: %w", i+1, err)
		}

		todo := todoItem{
			ID:        id,
			Title:     record[FieldIndexes["title"]],
			Completed: completed,
			Timestamp: timestamp,
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
