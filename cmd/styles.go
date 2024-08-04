package cmd

import (
	"fmt"
	"sort"
	"strings"
	"time"

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

var subtleTextStyle = lipgloss.NewStyle().
	Italic(true).
	Faint(true)

var Title = lipgloss.NewStyle().
	Underline(true)

var checkedSymbolStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(CheckedSymbolColor))

func printFinishedItem(todo todoItem) {
	idStr := fmt.Sprintf("%3d ", todo.ID) // Right-justified with a width of 3
	fmt.Println(" " + checkedSymbolStyle.Render(CheckedSymbol) + subtleTextStyle.Render(idStr) + " " + checkedItemStyle.Render(todo.Title))
}

func printUnfinishedItem(todo todoItem) {
	idStr := fmt.Sprintf("%3d ", todo.ID) // Right-justified with a width of 3
	fmt.Println(" " + UncheckedSymbol + subtleTextStyle.Render(idStr) + " " + todo.Title)
}

func categorizeItemsByDay(items []todoItem) map[time.Time][]todoItem {
	categorized := make(map[time.Time][]todoItem)

	for _, item := range items {
		// Truncate the timestamp to remove time information, keeping only the date
		date := item.Timestamp.Truncate(24 * time.Hour)

		// Append the item to the slice for this date
		categorized[date] = append(categorized[date], item)
	}

	return categorized
}

func printItems(items []todoItem) {
	categorizedItemsMap := categorizeItemsByDay(items)

	// Sort the days
	var days []time.Time
	for day := range categorizedItemsMap {
		days = append(days, day)
	}
	sort.Slice(days, func(i, j int) bool {
		return days[i].Before(days[j])
	})

	for _, day := range days {
		categorizedItems := categorizedItemsMap[day]

		// Format the day as a string and relative string
		dayStr := subtleTextStyle.Render(day.Format("Monday, January 2"))
		relativeDayStr := Title.Render(formatRelativeDate(day))

		indent := strings.Repeat(" ", 1)
		fmt.Printf("\n%s%s %s\n\n", indent, relativeDayStr, dayStr)

		for _, item := range categorizedItems {
			if item.Completed {
				printFinishedItem(item)
			} else {
				printUnfinishedItem(item)
			}
		}
	}
}
