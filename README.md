# Shell from scratch

A simple and minimalist shell written in Go.

## Features

- **External Command Execution**: Seamlessly run standard system commands.
- **Built-in Commands**:
  - `cd`: Change the current working directory.
  - `echo`: Print text to the terminal.
  - `pwd`: Print the absolute path to the current working directory.
  - `type`: Check how the shell interprets a specific command name (built-in, executable, etc.).
  - `exit`: Gracefully terminate the shell session.

## Building the Project

This project uses a `Makefile` to automate the build process across multiple operating systems. The binaries are automatically named using the latest Git tag for versioning.

To run the shell directly without building a binary:
```bash
make run
```

To build compiled binaries for Windows, macOS, and Linux:
```bash
make build
```

The compiled binaries will be output to the `build/` directory (e.g., `build/shell-v1.0.0-linux-amd64`).

## Usage

After compiling the project, you can run the binary directly from your terminal:

```bash
# Example for macOS Apple Silicon
./build/shell-latest-darwin-arm64
```