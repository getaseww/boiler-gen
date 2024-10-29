# Boiler-Gen

**Boiler-Gen** is a simple CLI tool designed to quickly generate boilerplate code files for your project, including models, repositories, services, controllers, and routes. This tool helps streamline project setup and standardize file structure for faster development.

## Features

- Generate standalone components like models, controllers, repositories, and services.
- Create comprehensive features that include model, repository, service, controller, and route files.
- Autocomplete support for faster command execution in shells.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (1.18 or newer) installed on your machine.

### Installation Command

To install `boiler-gen` globally, run:

```bash
go install github.com/getaseww/boiler-gen@latest
```

This will install the `boiler-gen` command to your `GOPATH/bin` directory, making it available globally.

Ensure that `GOPATH/bin` is in your PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## Usage

After installation, you can run `boiler-gen` commands to generate specific boilerplate files.

### Command Syntax

```bash
boiler-gen [command]
```

### Available Commands

- `completion`      - Generate the autocompletion script for the specified shell.
- `help`            - Show help for any command.
- `make:controller` - Generate a controller file.
- `make:feature`    - Generate model, repository, service, controller, and route files for a complete feature.
- `make:model`      - Generate a model file.
- `make:repository` - Generate a repository file.
- `make:route`      - Generate a route file.
- `make:service`    - Generate a service file.

### Example Usage

Generate a controller file:

```bash
boiler-gen make:controller
```

Generate a complete feature with model, repository, service, controller, and route files:

```bash
boiler-gen make:feature
```

To see details on using a command, add `--help` after the command:

```bash
boiler-gen make:model --help
```

### Flags

- `-h, --help`   - Show general help for `boiler-gen`.

## Autocompletion

To enable autocompletion, run the following command to generate the completion script for your preferred shell (e.g., bash, zsh):

```bash
boiler-gen completion > /path/to/completion/script
```

Follow the instructions in your shell's documentation to source the generated script.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
