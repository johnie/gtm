package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	errorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true)
	promptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Bold(true)
	infoStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#0000FF")).Bold(true)
	checkStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Bold(true)
)

func PrintError(message string) {
	fmt.Println(errorStyle.Render(message))
}

func PrintPrompt(message string) {
	fmt.Print(promptStyle.Render(message))
}

func PrintInfo(message string) {
	fmt.Println(infoStyle.Render(message))
}

func PrePrendCheckmark(message string) string {
	return checkStyle.Render("✔") + " " + message
}

func PrePrendError(message string) string {
	return errorStyle.Render("✖ ") + message
}

func TicketStyle(ticket string) string {
	return promptStyle.
		Render(ticket)
}
