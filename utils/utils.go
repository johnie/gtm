package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

func IsRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(output)) == "true"
}

func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func ExtractTicket(branchName string) (string, error) {
	re := regexp.MustCompile(`[A-Z]{2,}-\d+`)
	match := re.FindString(branchName)
	if match == "" {
		return "", fmt.Errorf("no ticket found in branch name")
	}
	return match, nil
}

func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func Copy(text string) error {
	switch runtime.GOOS {
	case "darwin":
		return copyDarwin(text)
	case "windows":
		return copyWindows(text)
	case "linux":
		return copyLinux(text)
	default:
		return errors.New("unsupported platform")
	}
}

func copyDarwin(text string) error {
	return runCommand("pbcopy", text)
}

func copyWindows(text string) error {
	return runCommand("clip", text)
}

func copyLinux(text string) error {
	return runCommand("xclip", "-selection", "clipboard", text)
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(args[0])); err != nil {
		return err
	}
	in.Close()
	return cmd.Wait()
}
