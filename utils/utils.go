package utils

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

type Utils struct{}

func NewUtils() *Utils {
	return &Utils{}
}

func (u *Utils) IsRepo() bool {
	_, err := os.Stat(".git")
	return err == nil
}

func (u *Utils) GetCurrentBranch() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func (u *Utils) ExtractTicket(branchName string) (string, error) {
	if branchName == "" {
		return "", fmt.Errorf("branch name cannot be empty")
	}

	re := regexp.MustCompile(`[A-Z]{2,}-\d+`)
	match := re.FindString(branchName)
	if match == "" {
		return "", fmt.Errorf("no ticket found in branch name: %s", branchName)
	}
	return match, nil
}

func (u *Utils) Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (u *Utils) Copy(text string) error {
	switch runtime.GOOS {
	case "darwin":
		return u.copyDarwin(text)
	case "windows":
		return u.copyWindows(text)
	case "linux":
		return u.copyLinux(text)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

func (u *Utils) copyDarwin(text string) error {
	return u.runCommand("pbcopy", text)
}

func (u *Utils) copyWindows(text string) error {
	return u.runCommand("clip", text)
}

func (u *Utils) copyLinux(text string) error {
	return u.runCommand("xclip", "-selection", "clipboard", text)
}

func (u *Utils) runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdin pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}

	if _, err := stdin.Write([]byte(args[0])); err != nil {
		return fmt.Errorf("failed to write to stdin: %w", err)
	}
	stdin.Close()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("command failed: %w", err)
	}

	return nil
}
