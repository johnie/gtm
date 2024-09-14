package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/johnie/gtm/lib/ui"
	"github.com/johnie/gtm/utils"
	"github.com/spf13/cobra"
)

var (
	version     bool
	copyFlag    bool
	messageFlag string
)

var rootCmd = &cobra.Command{
	Use:   "gtm",
	Short: "A helper tool for git commits with ticket prepended",
	RunE:  run,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print the version")
	rootCmd.Flags().BoolVarP(&copyFlag, "copy", "c", false, "Only copy the ticket value")
	rootCmd.Flags().StringVarP(&messageFlag, "message", "m", "", "Commit message")
}

func run(cmd *cobra.Command, args []string) error {
	if version {
		fmt.Println("gtm v1.0.0")
		return nil
	}

	if !utils.IsRepo() {
		return fmt.Errorf("this script can only be run inside a git repository")
	}

	branchName, err := utils.GetCurrentBranch()
	if err != nil {
		return fmt.Errorf("error getting current branch: %w", err)
	}

	ticket, err := utils.ExtractTicket(branchName)
	if err != nil {
		ui.PrintPrompt("Enter JIRA ticket (e.g., XXX-000): ")
		reader := bufio.NewReader(os.Stdin)
		ticketInput, _ := reader.ReadString('\n')
		ticket = strings.TrimSpace(ticketInput)
	}

	if copyFlag {
		if err := utils.Copy(ticket); err != nil {
			return fmt.Errorf("error copying to clipboard: %w", err)
		}
		fmt.Println(ui.PrePrendCheckmark(ui.TicketStyle(ticket)))
		return nil
	}

	commitMessage := getCommitMessage(args, ticket)
	fullMessage := fmt.Sprintf("%s: %s", ticket, commitMessage)

	if err := utils.Commit(fullMessage); err != nil {
		return fmt.Errorf("error running git commit: %w", err)
	}

	return nil
}

func getCommitMessage(args []string, ticket string) string {
	if messageFlag != "" {
		return messageFlag
	}
	if len(args) > 0 {
		return strings.Join(args, " ")
	}
	ui.PrintPrompt(ticket + ": ")
	reader := bufio.NewReader(os.Stdin)
	messageInput, _ := reader.ReadString('\n')
	return strings.TrimSpace(messageInput)
}
