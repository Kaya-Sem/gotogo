/*
Copyright © 2024 Kaya-Sem Van Cauwenberghe kayasemvc@gmail.com
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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotogo",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		purgeCompletedItems()
		if len(args) == 0 {
			printTodos()
		} else {
			createTodo(args[0])
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gotogo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printTodos() {
	filePath, err := getFilePath()

	if err != nil {
		fmt.Errorf("Could not get file path: %w", err)
		return
	}

	todoItems, err := readTodos(filePath)
	printItems(todoItems)
}

func createTodo(todo string) {
	filePath, err := getFilePath()
	if err != nil {
		fmt.Printf("Could not get file path: %v\n", err)
		return
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Could not open file: %v\n", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Could not read CSV file: %v\n", err)
		return
	}

	// Get the last ID
	var lastID int
	if len(records) > 0 {
		lastRecord := records[len(records)-1]
		lastID, err = strconv.Atoi(lastRecord[0])
		if err != nil {
			fmt.Printf("Could not parse last ID: %v\n", err)
			return
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

	// Append the new record to the file
	writer := csv.NewWriter(file)
	err = writer.Write(newRecord)
	if err != nil {
		fmt.Printf("Could not write to CSV file: %v\n", err)
		return
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Printf("Could not flush writer: %v\n", err)
		return
	}

	fmt.Printf("Todo item with id %d added successfully", newID)
}
