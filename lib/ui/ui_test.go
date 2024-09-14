package ui

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestPrintError(t *testing.T) {
	output := captureOutput(func() {
		PrintError("Test error message")
	})
	expected := errorStyle.Render("Test error message") + "\n"
	if output != expected {
		t.Errorf("PrintError output mismatch. Got: %q, Want: %q", output, expected)
	}
}

func TestPrintPrompt(t *testing.T) {
	output := captureOutput(func() {
		PrintPrompt("Test prompt message")
	})
	expected := promptStyle.Render("Test prompt message")
	if output != expected {
		t.Errorf("PrintPrompt output mismatch. Got: %q, Want: %q", output, expected)
	}
}

func TestPrintInfo(t *testing.T) {
	output := captureOutput(func() {
		PrintInfo("Test info message")
	})
	expected := infoStyle.Render("Test info message") + "\n"
	if output != expected {
		t.Errorf("PrintInfo output mismatch. Got: %q, Want: %q", output, expected)
	}
}

func TestPrePrendCheckmark(t *testing.T) {
	result := PrePrendCheckmark("Test message")
	expected := promptStyle.Render("✔") + " Test message"
	if result != expected {
		t.Errorf("PrePrendCheckmark output mismatch. Got: %q, Want: %q", result, expected)
	}
}

func TestPrePrendError(t *testing.T) {
	result := PrePrendError("Test message")
	expected := errorStyle.Render("✖ ") + "Test message"
	if result != expected {
		t.Errorf("PrePrendError output mismatch. Got: %q, Want: %q", result, expected)
	}
}

func TestTicketStyle(t *testing.T) {
	result := TicketStyle("ABC-123")
	expected := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("5")).Render("ABC-123")
	if result != expected {
		t.Errorf("TicketStyle output mismatch. Got: %q, Want: %q", result, expected)
	}
}
