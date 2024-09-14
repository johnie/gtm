package utils

import (
	"os"
	"os/exec"
	"testing"
)

func TestIsRepo(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(currentDir)

	tempDir, err := os.MkdirTemp("", "test-repo")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	if IsRepo() {
		t.Error("Expected false for non-repo directory")
	}

	cmd := exec.Command("git", "init")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to initialize git repo: %v", err)
	}

	if !IsRepo() {
		t.Error("Expected true for repo directory")
	}
}

func TestGetCurrentBranch(t *testing.T) {
	branch, err := GetCurrentBranch()
	if err != nil {
		t.Fatalf("Failed to get current branch: %v", err)
	}
	if branch == "" {
		t.Error("Expected non-empty branch name")
	}
}

func TestExtractTicket(t *testing.T) {
	testCases := []struct {
		branchName string
		expected   string
		shouldErr  bool
	}{
		{"feature/ABC-123-add-new-feature", "ABC-123", false},
		{"bugfix/XY-789-fix-critical-bug", "XY-789", false},
		{"develop", "", true},
		{"feature/add-new-feature", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.branchName, func(t *testing.T) {
			ticket, err := ExtractTicket(tc.branchName)
			if tc.shouldErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if ticket != tc.expected {
					t.Errorf("Expected ticket %s, got %s", tc.expected, ticket)
				}
			}
		})
	}
}

func TestCommit(t *testing.T) {
	err := Commit("Test commit message")
	if err != nil {
		t.Errorf("Commit failed: %v", err)
	}
}

func TestCopy(t *testing.T) {
	err := Copy("Test text")
	if err != nil {
		t.Errorf("Copy failed: %v", err)
	}
}
