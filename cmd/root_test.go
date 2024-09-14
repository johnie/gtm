package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/johnie/gtm/utils"
	"github.com/spf13/cobra"
)

var isRepo = utils.IsRepo
var getCurrentBranch = utils.GetCurrentBranch
var commit = utils.Commit

func mockIsRepo() bool {
	return true
}

func mockGetCurrentBranch() (string, error) {
	return "feature/ABC-123-test-branch", nil
}

func mockCommit(message string) error {
	return nil
}

func TestRun(t *testing.T) {
	isRepo = mockIsRepo
	getCurrentBranch = mockGetCurrentBranch
	commit = mockCommit

	defer func() {
		isRepo = utils.IsRepo
		getCurrentBranch = utils.GetCurrentBranch
		commit = utils.Commit
	}()

	testCases := []struct {
		name        string
		args        []string
		messageFlag string
		copyFlag    bool
		wantErr     bool
	}{
		{
			name:        "Commit with args",
			args:        []string{"Test", "commit", "message"},
			messageFlag: "",
			copyFlag:    false,
			wantErr:     false,
		},
		{
			name:        "Commit with message flag",
			args:        []string{},
			messageFlag: "Test commit message",
			copyFlag:    false,
			wantErr:     false,
		},
		{
			name:        "Copy ticket",
			args:        []string{},
			messageFlag: "",
			copyFlag:    true,
			wantErr:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			cmd := &cobra.Command{}
			messageFlag = tc.messageFlag
			copyFlag = tc.copyFlag

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := run(cmd, tc.args)

			w.Close()
			os.Stdout = oldStdout

			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := buf.String()

			if (err != nil) != tc.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tc.wantErr)
			}

			if tc.copyFlag && !bytes.Contains([]byte(output), []byte("Copied ticket")) {
				t.Errorf("expected output to contain 'Copied ticket', got %s", output)
			}
		})
	}
}
