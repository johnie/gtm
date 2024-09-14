package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/johnie/gtm/lib/ui"
	"github.com/johnie/gtm/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	versionFlag bool
	copyFlag    bool
	messageFlag string
	urlFlag     bool
)

var rootCmd = &cobra.Command{
	Use:   "gtm",
	Short: "A helper tool for git commits with ticket prepended",
	RunE:  run,
}

var Utils = utils.NewUtils()

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	loadConfig()

	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Print the version")
	rootCmd.Flags().BoolVarP(&copyFlag, "copy", "c", false, "Only copy the ticket value")
	rootCmd.Flags().StringVarP(&messageFlag, "message", "m", "", "Commit message")
	rootCmd.Flags().BoolVarP(&urlFlag, "url", "u", false, "Print Ticket URL")
}

func loadConfig() {
	viper.SetConfigName(".gtmconfig")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME")

	viper.SetDefault("core.jira_url", "")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Could not get home directory: %v\n", err)
		return
	}

	configPath := filepath.Join(homeDir, ".gtmconfig")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		createDefaultConfig(configPath)
	}

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Could not read config file: %v\n", err)
	}
}

func createDefaultConfig(configPath string) {
	defaultConfig := []byte(`[core]
jira_url = "https://jira.example.com/browse/"
`)
	err := os.WriteFile(configPath, defaultConfig, 0644)
	if err != nil {
		fmt.Printf("Could not create default config file: %v\n", err)
	} else {
		fmt.Printf("Default config file created at %s\n", configPath)
	}
}

func run(cmd *cobra.Command, args []string) error {
	if versionFlag {
		fmt.Println("gtm v1.0.0")
		return nil
	}

	if !Utils.IsRepo() {
		return fmt.Errorf("this script can only be run inside a git repository")
	}

	branchName, err := Utils.GetCurrentBranch()
	if err != nil {
		return fmt.Errorf("error getting current branch: %w", err)
	}

	ticket, err := Utils.ExtractTicket(branchName)
	if err != nil {
		ui.PrintPrompt("Enter JIRA ticket (e.g., XXX-000): ")
		reader := bufio.NewReader(os.Stdin)
		ticketInput, _ := reader.ReadString('\n')
		ticket = strings.TrimSpace(ticketInput)
	}

	if urlFlag {
		jiraURL := viper.GetString("core.jira_url")
		if jiraURL == "" {
			return fmt.Errorf("JIRA URL not set in config")
		}
		if copyFlag {
			if err := Utils.Copy(jiraURL + ticket); err != nil {
				return fmt.Errorf("error copying to clipboard: %w", err)
			}
		}
		fmt.Println(ui.PrePrendCheckmark(ui.TicketStyle(jiraURL + ticket)))
		return nil
	}

	if copyFlag {
		if err := Utils.Copy(ticket); err != nil {
			return fmt.Errorf("error copying to clipboard: %w", err)
		}
		fmt.Println(ui.PrePrendCheckmark(ui.TicketStyle(ticket)))
		return nil
	}

	commitMessage := getCommitMessage(args, ticket)
	fullMessage := fmt.Sprintf("%s: %s", ticket, commitMessage)

	if err := Utils.Commit(fullMessage); err != nil {
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
