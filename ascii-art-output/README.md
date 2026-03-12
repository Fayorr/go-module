# ASCII Art Output

A Go program that generates ASCII art from text input and writes the output to a file. The program uses font files (banner files) to convert user-provided text into ASCII art representation and saves the result to a specified output file.

## Features

- Generate ASCII art for text input
- Write ASCII art output to a file with `--output=` flag
- Support for multi-line text using `\\n` separator
- Multiple banner styles: `standard`, `shadow`, `thinkertoy`
- Handles special characters and spaces
- Font-based rendering (8 lines per character)

## Installation

1. Ensure you have Go installed (version 1.16 or later).
2. Clone or download the project.
3. Run `go mod tidy` to download dependencies (if any).

## Usage

Build and run the program:

```bash
go build -o ascii-art-output
./ascii-art-output --output=banner.txt "Hello World" standard
```

Or run directly:

```bash
go run . --output=banner.txt "Hello World" standard
```

### Arguments

1. `--output=<filename>`: Flag specifying the output file path (required)
2. `<text>`: The text to convert to ASCII art (required)
3. `<banner>`: The font style to use - `standard`, `shadow`, or `thinkertoy` (required)

### Examples

- Basic: `go run . --output=result.txt "Hello" standard`
- Multi-line: `go run . --output=result.txt "Hello\\nWorld" shadow`
- With special characters: `go run . --output=result.txt "Hi {There}" thinkertoy`

### Output

The program writes ASCII art to the specified file. For example:

```bash
go run . --output=banner.txt "Hello" standard
```

Creates `banner.txt` with ASCII art output.

## Project Structure

- `main.go`: Entry point, argument validation, and file output handling
- `runner.go`: Core logic for generating ASCII art
- `runner_test.go`: Unit tests with examples
- `banner/`: Directory containing font files
  - `standard.txt`: Classic ASCII art font
  - `shadow.txt`: Shadowed ASCII art font
  - `thinkertoy.txt`: Alternative ASCII art font
- `test-files/`: Example test input files

## Supported Banners

- **standard**: Classic ASCII art font
- **shadow**: Darker, shadowed ASCII art font
- **thinkertoy**: Alternative ASCII art font style

## Error Handling

The program validates input and reports errors:
- Checks for exactly 3 arguments
- Verifies the `--output=` flag is present
- Reports file write errors with appropriate messages

## License

This project is for educational purposes.
