# ASCII Art File System (ASCII-Art-FS)

Convert text into stylized ASCII art using file system-based banner fonts. This Go project reads banner definition files to render any input text in three distinct ASCII art styles.

## Features

- **Multiple Banner Styles**: Choose from three ASCII art fonts:
  - `standard` - Classic block letters
  - `shadow` - Shaded/3D effect letters
  - `thinkertoy` - Playful geometric letters
- **Multi-line Support**: Use `\n` escape sequences to create multi-line ASCII art
- **Simple Command Interface**: Intuitive CLI with sensible defaults
- **File System Based**: Banner definitions stored as readable text files

## Installation

### Prerequisites
- Go 1.25.0 or higher

### Setup
```bash
git clone <repository-url>
cd ascii-art-fs
go mod download
```

## Usage

### Basic Syntax
```bash
go run . [STRING] [BANNER]
```

### Arguments
- `STRING` (required): The text to convert to ASCII art
- `BANNER` (optional): The banner style to use (`standard`, `shadow`, or `thinkertoy`). Defaults to `standard` if omitted.

### Examples

**Basic Usage (Default Banner)**
```bash
go run . "Hello"
```

**With Specific Banner**
```bash
go run . "Hello" standard
go run . "Hello" shadow
go run . "Hello" thinkertoy
```

**Multi-line ASCII Art**
```bash
go run . "Hello\nWorld" standard
```

## Project Structure

```
ascii-art-fs/
├── main.go           # Entry point with input validation
├── runner.go         # Core ASCII art generation logic
├── runner_test.go    # Test suite
├── go.mod            # Module definition
├── README.md         # This file
├── banner/           # Banner definition files
│   ├── standard.txt  # Standard ASCII art font
│   ├── shadow.txt    # Shadow ASCII art font
│   └── thinkertoy.txt # Thinkertoy ASCII art font
├── test-files/       # Expected output files for testing
│   ├── audit1.txt - audit9.txt
└── hello world/      # Example output directory
```

## How It Works

### Banner File Format
Each banner file contains ASCII art definitions for all printable characters (space through tilde). Characters are arranged in 9-line blocks:
- Line 0: Blank separator
- Lines 1-8: Character definition (8 rows per character)

### Rendering Process
1. **Parse Input**: Split the input text by `\n` escape sequences to handle multiple lines
2. **Load Banner**: Read the appropriate banner definition file
3. **Generate Output**: For each line in the input:
   - For each of the 8 rows in a character block:
     - Iterate through each character and extract its corresponding ASCII art line
     - Concatenate horizontally to form complete rows
   - Output each row on a new line

## Testing

Run the test suite to verify functionality:
```bash
go test -v
```

The test suite includes 9 audit cases validating various input scenarios with different banners and multi-line handling.

## Technical Notes

- ASCII characters are stored sequentially in banner files using a consistent offset calculation: `position = (charCode - 32) * 9`
- Characters are rendered 8 rows tall
- Empty lines in input (`\n` sequences) produce blank lines in output
- File I/O errors gracefully exit with error messages

## Educational Context

This project is part of an educational curriculum focused on:
- File I/O operations in Go
- String manipulation and parsing
- Algorithm design for text rendering
- Test-driven development practices

## License

Educational project

## Author

Created as part of the 01-edu curriculum
