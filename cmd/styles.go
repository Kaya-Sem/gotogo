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

var checkedSymbolStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(CheckedSymbolColor))

func printFinishedItem(todo todoItem) {
	fmt.Println(checkedSymbolStyle.Render(CheckedSymbol) + checkedItemStyle.Render(todo.Name))
}

func printUnfinishedItem(todo todoItem) {
	fmt.Println(UncheckedSymbol + todo.Name)
}
