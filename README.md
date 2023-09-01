# warestack-cli

A CLI tool built using the Cobra framework for managing Warestack resources. It supports authentication through Firebase
and allows users to manage their workspaces, services and Cloud deployments. It also supports local development and
applying changes to their services.

## Project Structure

```css
├── cmd
│   ├── root.go
│   ├── configure
│   │   ├── configure.go
│   │   ├── current.go
│   │   ├── list.go
│   │   └── set.go
│   └── login
│       ├── login.go
│       └── server.go
├── pkg
│   ├── api
│   │   ├── client.go
│   │   ├── user.go
│   │   └── organization.go
│   ├── auth
│   │   ├── auth.go
│   │   ├── browser.go
│   │   └── redirect.go
│   ├── ui
│   │   └── logger.go
│   └── util
│       ├── file.go
│       └── paths.go
├── main.go
└── README.md
```

## Packages

- **cmd**: Contains root and CLI commands, structured with Cobra.
  - **configure**: Commands to manage active organization (`current`, `list`, `set`). 
  - **login**: Commands and server logic for Firebase-based login (`login.go`, `server.go`).
- **pkg**: Contains the main packages that the CLI tool uses. 
  - **api**: Functions for interacting with the remote API (`client.go`, `user.go`, `organization.go`). 
  - **auth**: Authentication handling, including browser-based Firebase auth (`auth.go`, `browser.go`, `redirect.go`). 
  - **ui**: User interface helpers (`logger.go`). 
  - **util**: Utility functions (`file.go` for JSON file operations and `paths.go` for path management).

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/warestack/warestack-cli.git
    ```

2. Compile the project:

    ```bash
    go build
    ```

3. Add the `warestack-cli` binary to your system's PATH:

   - For Linux and macOS, move the `warestack-cli` binary to `/usr/local/bin`:
   
    ```bash
    sudo mv warestack /usr/local/bin
    ```

    - For Windows, move the `warestack-cli.exe` binary to a folder in your system's PATH or add the folder containing 
      `warestack-cli.exe` to your system's PATH environment variable.

## Usage

The warestack-cli tool has two main commands: `login` and `configure`.

```bash
# Basic usage
warestack-cli [command]

# To be authenticated
warestack-cli auth

# To configure the active organization
warestack-cli configure
```

_Note that tokens and configurations are stored in local JSON files located in the user's home directory under a hidden
folder named .warestack._

### License

This project is licensed under the terms of the MIT license.

