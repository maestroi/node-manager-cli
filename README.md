
# Node Manager CLI

Node Manager CLI is a command-line tool designed to manage the setup, update, and maintenance of Nimiq nodes. It simplifies the process of deploying, configuring, and updating Nimiq nodes using Ansible.

## Features

- **Setup Nimiq Node**: Easily set up a Nimiq node for various networks and node types.
- **Full Stack or Bare Install**: Default setup includes nginx, watchdog, and monitoring; homelab mode installs only the node.
- **Update Nimiq Node**: Check for updates and update the Nimiq node to the latest version.
- **Cleanup**: Remove all configurations and files related to the Nimiq node setup.
- **List Supported Configurations**: View all supported protocols, networks, and node types.
- **Version Management**: View the CLI version.

## Easy install

Install the latest release to `/usr/local/bin`:

```sh
curl -fsSL https://raw.githubusercontent.com/maestroi/node-manager-cli/main/scripts/install.sh | sudo bash
```

Install a specific version:

```sh
curl -fsSL https://raw.githubusercontent.com/maestroi/node-manager-cli/main/scripts/install.sh | sudo VERSION=v0.8.0 bash
```

Download only (without installing to `/usr/local/bin`):

```sh
curl -L -o node-manager-cli https://github.com/maestroi/node-manager-cli/releases/latest/download/node-manager-cli && chmod +x node-manager-cli
```

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/node-manager-cli.git
    cd node-manager-cli
    ```

2. Build the CLI (set the version to match your release tag):

    ```sh
    go build -ldflags "-X node-manager-cli/config.CLIVersion=v1.0.0" -o node-manager-cli
    ```

3. (Optional) Move the binary to `/usr/local/bin` for global access:

    ```sh
    sudo mv node-manager-cli /usr/local/bin/
    ```

## Usage

### Setup

Set up a Nimiq node:

```sh
sudo ./node-manager-cli setup --network <network> --node-type <node-type> --protocol <protocol>
```

- `<network>`: The network to deploy the node on (default: `testnet`).
- `<node-type>`: The type of the node (`validator`, `full_node`, `history_node`).
- `<protocol>`: The protocol to deploy (`nimiq`).
- `<branch>`: The branch to use for the ansible repository (e.g., `main`).
- `--no-monitor`: Skip Grafana, Prometheus, Loki, and related monitoring containers.
- `--homelab`: Bare install with only the Nimiq node container.

#### Full stack (default)

Installs the node plus nginx, watchdog, validator activator (for validators), and monitoring:

```sh
sudo node-manager-cli setup --network testnet --node-type validator --protocol nimiq
```

Skip monitoring only:

```sh
sudo node-manager-cli setup --network testnet --node-type validator --no-monitor
```

#### Bare homelab install

For a homelab setup where you expose RPC directly and manage monitoring yourself:

```sh
sudo node-manager-cli setup --network mainnet --node-type validator --homelab
```

This installs only the Nimiq node container and exposes:

- RPC on `0.0.0.0:8648`
- P2P on `0.0.0.0:8443`

Access RPC from your LAN, for example:

```sh
curl http://192.168.1.100:8648
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
- [pre-commit](https://pre-commit.com/) (optional, for local hooks)

### Pre-commit

Install hooks once:

```sh
pip install pre-commit
pre-commit install
```

Run on all files:

```sh
pre-commit run --all-files
```

The same checks run in CI on pull requests and pushes to `main` via GitHub Actions.

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

Set `CLIVersion` at build time so `version` and `autoupdate` report the correct release:

```sh
go build -ldflags "-X node-manager-cli/config.CLIVersion=v1.0.0" -o node-manager-cli
```

Without `-ldflags`, the default in `config/version.go` is used (`1.0.0`). GitHub release builds set this from the release tag automatically.

### Running Tests

To run the tests, use:

```sh
go test ./...
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
