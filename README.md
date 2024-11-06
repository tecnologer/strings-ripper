## Ripper Tool

Ripper is a command-line tool designed to split a given input string into multiple lines, each with a specified maximum length. This can be particularly useful for formatting long strings to fit within certain size constraints.

### Features

- Split input strings into multiple lines based on a specified length.
- Automatically handle input from both command-line arguments and standard input (pipe).

### Usage

1. **Installation**: Ensure that Go is installed on your system and that the GOPATH is set. Then, compile the tool using `go build` in the root directory of the project.

2. **Command-Line Arguments**:
   - `-i` or `--input`: Specify the input string directly.
   - `-l` or `--length`: Specify the maximum length of each line in the output. Default is 150 characters.

3. **Examples**:
   - Direct input:
     ```
     ./ripper -i "This is a very long string that needs to be split into smaller chunks." -l 10
     ```
   - Using pipe:
     ```
     echo "This is a very long string that needs to be split into smaller chunks." | ./ripper -l 10
     ```

The output will be the input string split into lines, each not exceeding the specified length, formatted and concatenated with a '+' at the end of each line except the last.
