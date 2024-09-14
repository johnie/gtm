# GTM - Git Ticket Manager

GTM is a CLI tool that assists with Git commit operations by prepending the current branch's ticket ID to commit messages. It also provides functionalities to copy ticket IDs to the clipboard.

## Features

- Auto-prepend JIRA ticket IDs from branch names to Git commit messages.
- Copy ticket IDs to the clipboard.
- Customizable commit messages via command-line flags.
- Supports multiple operating systems (macOS, Windows, Linux).

## Installation

```console
 $ curl -fsSL https://raw.githubusercontent.com/johnie/gtm/main/bin/install.sh | sh
```

All this installation script does is download the `gtm` script, make it an executable, and copy it to your `$PATH (/usr/local/bin)`. For copying to your `$PATH`, it may be required to enter your password. If there is a better way to do this, please send in a pull request.

## Configuration

GTM uses a configuration file named `.gtmconfig` to store settings such as the JIRA base URL. This configuration file is located in your home directory.

### Default Configuration

If a `.gtmconfig` file does not exist, GTM will create a default configuration file with the following content:

```toml
[core]
jira_url = "https://jira.example.com/browse/"
```

### Custom Configuration

You can customize the configuration by editing the `.gtmconfig` file in your home directory. For example, to set your JIRA URL, modify the `jira_url` field:

```toml
[core]
jira_url = "https://yourcompany.jira.com/browse/"
```

## Usage

### Basic Usage

To run gtm with default configurations:

```sh
gtm
```

### Command-Line Flags

- `-c, --copy`: Only copy the ticket value to the clipboard.
- `-m, --message`: Custom commit message.
- `-u, --url`: Print the JIRA ticket URL.
- `-v, --version`: Print the GTM version

### Examples

1. Committing with a prompt: `gtm`
2. Committing with a message provided via a flag: `gtm -m "Initial commit"`
3. Copying the ticket value to the clipboard: `gtm -c`

## Development

This section provides an overview of the codebase and how to contribute to GTM.

### Directory Structure

```
.
├── cmd – Contains the command logic for the GTM
│   ├── root.go
│   └── root_test.go
├── lib
│   └── ui – Provides utilities for styling command-line outputs.
│       ├── ui.go
│       └── ui_test.go
├── utils – Includes helper functions for interacting with Git.
│   └── utils.go
│   └── utils_test.go
├── go.mod
├── go.sum
├── main.go
```

### Running Tests

To run the tests, use the following command:

```sh
go test ./...
```

## Contributing

We welcome contributions to the project! To contribute:

1. Fork the repository.
2. Create a new branch with a descriptive name.
3. Implement your changes.
4. Ensure all tests pass.
5. Submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

## Acknowledgements

- [Cobra https://github.com/spf13/cobra](https://github.com/spf13/cobra): For command-line interface management.
- [Lipgloss https://github.com/charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss): For styling command-line output.

---

Thank you for using GTM! Feel free to submit issues or feature requests on the GitHub repository https://github.com/johnie/gtm.
