# ASCII Art

A Go program that converts text into ASCII art using a standard character set.

## Description

This project takes command-line input and renders it as ASCII art by reading character patterns from a `standard.txt` file. Each character is mapped to an 8-line ASCII representation, allowing you to display text in a large, visually distinctive format.

## Features

- Convert text to ASCII art
- Support for multiple words (joins them with spaces)
- Support for line breaks using `\n` escape sequence
- Reads character patterns from `standard.txt`

## Requirements

- Go 1.25.0 or later

## Usage

```bash
go run main.go runner.go "your text here"
```

### Examples

```bash
# Display a single word
go run main.go runner.go "Hello"

# Display multiple words
go run main.go runner.go "Hello World"

# Display multiple lines (use \n for line breaks)
go run main.go runner.go "Hello\nWorld"
```

## Project Structure

- `main.go` - Entry point that validates input and manages file reading
- `runner.go` - Core logic that converts text to ASCII art
- `standard.txt` - Character pattern definitions for ASCII rendering
- `runner_test.go` - Unit tests for the ASCII art functionality
- `go.mod` - Go module definition

## How It Works

1. The program reads ASCII art character patterns from `standard.txt`
2. Each character occupies 9 positions in the file (1 for the character itself, 8 for its ASCII representation)
3. For each input character, the program looks up its pattern and retrieves 8 lines
4. All lines at the same height are combined to form the output

## Testing

Run the test suite:

```bash
go test
```

## License

Educational project for ASCII art generation practice.
