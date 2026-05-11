# ASCII Art Web

## Description

ASCII Art Web is a web-based application that converts user input text into ASCII art using different banner styles. The application provides an interactive interface where users can select from multiple banner formats (standard, shadow, and thinkertoy) and generate beautifully formatted ASCII art output.

## Authors

- L2E Team

## Usage: How to Run

### Prerequisites

- Go 1.16 or higher

### Installation & Running

1. Navigate to the project directory:

```bash
cd ascii-art-web
```

2. Run the application:

```bash
go run .
```

3. Open your web browser and navigate to:

```
http://localhost:3000
```

4. Enter your text in the input field, select a banner style, and click submit to generate ASCII art.

## Implementation Details: Algorithm

### Architecture

The application follows a simple MVC-like pattern with:

- **Main Server** (`main.go`): Sets up HTTP routes and starts the server on port 3000
- **Handlers** (`handlers.go`): Manages HTTP requests and responses
- **Banner Files** (`banner/`): Contains ASCII art character definitions for different styles
- **Templates** (`templates/`): HTML templates for the user interface
- **Runner** (`runner.go`): Core logic for converting text to ASCII art

### ASCII Art Generation Algorithm

1. **Banner Selection**: The application reads the appropriate banner file based on user selection (standard.txt, shadow.txt, or thinkertoy.txt)

2. **Character Mapping**: Each printable ASCII character is represented in the banner file across 9 lines:
   - Each character occupies 9 lines in sequence
   - Characters are indexed by their ASCII value offset from space (32)
3. **Text Processing**:
   - Input text is split by newline characters
   - For each line in the input, the algorithm iterates through 8 visual lines (0-7)
   - For each character in the input line, it calculates the starting position: `pos = (ASCII_value - 32) * 9`
   - The corresponding ASCII art for that character at the current line is retrieved and concatenated

4. **Output Generation**: The generated ASCII art is formatted and returned to the user through the web interface

### Routes

- `GET /`: Displays the home page with input form
- `POST /ascii-art`: Processes form submission and returns generated ASCII art

### Banner Styles

- **standard.txt**: Classic ASCII art style
- **shadow.txt**: Shadow effect style
- **thinkertoy.txt**: Thin character style
