package utils

import (
	"testing"
)

type MockUtils struct {
	IsRepoFunc           func() bool
	GetCurrentBranchFunc func() (string, error)
	ExtractTicketFunc    func(string) (string, error)
	CommitFunc           func(string) error
	CopyFunc             func(string) error
}

func (m *MockUtils) IsRepo() bool {
	return m.IsRepoFunc()
}

func (m *MockUtils) GetCurrentBranch() (string, error) {
	return m.GetCurrentBranchFunc()
}

func (m *MockUtils) ExtractTicket(branchName string) (string, error) {
	return m.ExtractTicketFunc(branchName)
}

func (m *MockUtils) Commit(message string) error {
	return m.CommitFunc(message)
}

func (m *MockUtils) Copy(text string) error {
	return m.CopyFunc(text)
}

func TestCommit(t *testing.T) {
	mockUtils := &MockUtils{
		CommitFunc: func(message string) error {
			if message != "Test commit message" {
				t.Errorf("Expected 'Test commit message', got %s", message)
			}
			return nil
		},
	}

	err := mockUtils.Commit("Test commit message")
	if err != nil {
		t.Errorf("Commit failed: %v", err)
	}
}

func TestGetCurrentBranch(t *testing.T) {
	mockUtils := &MockUtils{
		GetCurrentBranchFunc: func() (string, error) {
			return "main", nil
		},
	}

	branch, err := mockUtils.GetCurrentBranch()
	if err != nil {
		t.Fatalf("Failed to get current branch: %v", err)
	}
	if branch != "main" {
		t.Errorf("Expected branch 'main', got %s", branch)
	}
}

func TestExtractTicket(t *testing.T) {
	mockUtils := &MockUtils{
		ExtractTicketFunc: func(branchName string) (string, error) {
			return "TEST-123", nil
		},
	}

	ticket, err := mockUtils.ExtractTicket("feature/TEST-123")
	if err != nil {
		t.Fatalf("Failed to extract ticket: %v", err)
	}
	if ticket != "TEST-123" {
		t.Errorf("Expected ticket 'TEST-123', got %s", ticket)
	}
}

func TestIsRepo(t *testing.T) {
	mockUtils := &MockUtils{
		IsRepoFunc: func() bool {
			return true
		},
	}

	if !mockUtils.IsRepo() {
		t.Error("Expected to be in a repo")
	}
}

func TestCopy(t *testing.T) {
	mockUtils := &MockUtils{
		CopyFunc: func(text string) error {
			if text != "Test text" {
				t.Errorf("Expected 'Test text', got %s", text)
			}
			return nil
		},
	}

	err := mockUtils.Copy("Test text")
	if err != nil {
		t.Errorf("Copy failed: %v", err)
	}
}
