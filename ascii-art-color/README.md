# ASCII Art Color

A simple Go program that generates ASCII art from text input using a predefined font. The program reads a font file (`standard.txt`) and converts user-provided text into ASCII art representation.

## Features

- Generate ASCII art for text input
- Support for multi-line text using `\\n` separator
- Handles special characters and spaces
- Font-based rendering (8 lines per character)

## Installation

1. Ensure you have Go installed (version 1.25.0 or later).
2. Clone or download the project.
3. Run `go mod tidy` to download dependencies (if any).

## Usage

Build and run the program:

```bash
go build -o ascii-art-color
./ascii-art-color "Hello World"
```

Or run directly:

```bash
go run . "Hello World"
```

### Examples

- Single line: `go run . "Hello"`
- Multi-line: `go run . "Hello\\nWorld"`
- With special characters: `go run . "{Hello & There #}"`

### Output

The program outputs ASCII art to the console. For example, input `"hello\\n\\nthere"` produces:

```
 _              _   _
| |            | | | |
| |__     ___  | | | |   ___
|  _ \   / _ \ | | | |  / _ \
| | | | |  __/ | | | | | (_) |
|_| |_|  \___| |_| |_|  \___/


 _     _
| |   | |
| |_  | |__     ___   _ __    ___
| __| |  _ \   / _ \ | '__|  / _ \
\ |_  | | | | |  __/ | |    |  __/
 \__| |_| |_|  \___| |_|     \___|
```

## Project Structure

- `main.go`: Entry point, argument validation, and file reading
- `runner.go`: Core logic for generating ASCII art
- `runner_test.go`: Unit tests with examples
- `standard.txt`: Font file containing ASCII art characters
- `test.txt`: Example output file

## Future Enhancements

- Color support for ASCII art output (ANSI color codes)
- Additional font options
- Command-line flags for customization

## License

This project is for educational purposes.