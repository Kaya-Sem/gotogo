/*
Copyright © 2024 Kaya-Sem Van Cauwenberghe kayasemvc@gmail.com
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotogo",
	Short: "A brief description of your application",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		purgeCompletedItems()
		printTodos()
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// TODO: export to more relavant file
func printTodos() {
	filePath, err := getFilePath()

	if err != nil {
		fmt.Errorf("Could not get file path: %w", err)
		return
	}

	todoItems, err := readTodos(filePath)
	printItems(todoItems)
}
