
# Node Manager CLI

Node Manager CLI is a command-line tool designed to manage the setup, update, and maintenance of Nimiq nodes. It simplifies the process of deploying, configuring, and updating Nimiq nodes using Ansible.

## Features

- **Setup Nimiq Node**: Easily set up a Nimiq node for various networks and node types.
- **Update Nimiq Node**: Check for updates and update the Nimiq node to the latest version.
- **Cleanup**: Remove all configurations and files related to the Nimiq node setup.
- **List Supported Configurations**: View all supported protocols, networks, and node types.
- **Version Management**: View the CLI version.

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/node-manager-cli.git
    cd node-manager-cli
    ```

2. Build the CLI:

    ```sh
    go build -o node-manager-cli
    ```

3. (Optional) Move the binary to `/usr/local/bin` for global access:

    ```sh
    sudo mv node-manager-cli /usr/local/bin/
    ```

## Usage

### Setup

Set up a Nimiq node:

```sh
sudo node-manager-cli setup --network <network> --node-type <node-type> --protocol <protocol> --branch <branch>
```

- `<network>`: The network to deploy the node on (default: `testnet`).
- `<node-type>`: The type of the node (`validator`, `full_node`, `history_node`).
- `<protocol>`: The protocol to deploy (`nimiq`).
- `<branch>`: The branch to use for the protocol repository (e.g., `master`, `main`).

Example:

```sh
sudo node-manager-cli setup --network testnet --node-type validator --protocol nimiq --branch master
```

### Update

Update the Nimiq node to the latest version:

```sh
sudo node-manager-cli update
```

Force update even if the latest version is already installed:

```sh
sudo node-manager-cli update --force
```

Specify a branch for the update:

```sh
sudo node-manager-cli update --branch <branch>
```

### Cleanup

Remove all configurations and files related to the Nimiq node setup:

```sh
sudo node-manager-cli cleanup
```

### List Supported Configurations

View all supported protocols, networks, and node types:

```sh
node-manager-cli list
```

### Version

View the CLI version:

```sh
node-manager-cli version
```

## Development

### Prerequisites

- Go 1.18+
- Git

### Project Structure

```
node-manager-cli/
├── cmd/
│   ├── cleanup.go
│   ├── list.go
│   ├── root.go
│   ├── setup.go
│   ├── update.go
│   ├── version.go
├── config/
│   ├── config.go
│   ├── constants.go
├── setup/
│   ├── config.go
│   ├── dependencies.go
│   ├── repository.go
│   ├── run.go
│   ├── utils.go
├── go.mod
├── go.sum
├── main.go
```

### Building the CLI

To build the CLI, run:

```sh
go build -o node-manager-cli
```

### Running Tests

To run the tests, use:

```sh
go test ./...
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
