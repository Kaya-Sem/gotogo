package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	CheckedSymbolColor = "#04B575" // green
	CheckedSymbol      = "✓"
	UncheckedSymbol    = " "
)

var checkedItemStyle = lipgloss.NewStyle().
	Italic(true).
	Faint(true).
	Strikethrough(true)

var uncheckedIDStyle = lipgloss.NewStyle().
	Italic(true).
	Faint(true)

var checkedIDStyle = lipgloss.NewStyle().
	Italic(true).
	Faint(true)

var checkedSymbolStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(CheckedSymbolColor))

func printFinishedItem(todo todoItem) {
	idStr := fmt.Sprintf("%3d ", todo.ID) // Right-justified with a width of 3
	fmt.Println(" " + checkedSymbolStyle.Render(CheckedSymbol) + checkedIDStyle.Render(idStr) + " " + checkedItemStyle.Render(todo.Title))
}

func printUnfinishedItem(todo todoItem) {
	idStr := fmt.Sprintf("%3d ", todo.ID) // Right-justified with a width of 3
	fmt.Println(" " + UncheckedSymbol + uncheckedIDStyle.Render(idStr) + " " + todo.Title)
}

func printItems(items []todoItem) {
	for _, item := range items {
		if item.Completed {
			printFinishedItem(item)
		} else {
			printUnfinishedItem(item)
		}
	}
}
