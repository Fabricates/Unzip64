# Unzip64

A simple Go command line tool that reads base64-encoded, flate-compressed content from a file or stdin, decodes it, decompresses it, and prints the result to stdout.

## Features

- Read base64 content from a specified file or from stdin
- Decode base64 content
- Decompress flate-compressed data
- Output decompressed content to stdout

## Installation

### Download pre-built binaries

Pre-built binaries are available from the [releases page](https://github.com/fabricates/unzip64/releases). Nightly releases are automatically built and published.

### Build from source

```bash
git clone https://github.com/fabricates/unzip64.git
cd unzip64
make build
```

Or without Make:

```bash
go build -o unzip64 .
```

### Install using go install

```bash
go install github.com/fabricates/unzip64@latest
```

## Usage

### Read from a file

```bash
./unzip64 input.txt
```

### Read from stdin

```bash
echo "base64-encoded-flate-content" | ./unzip64
```

```bash
cat input.txt | ./unzip64
```

## Examples

### Example 1: Process a file containing base64-encoded flate data

```bash
./unzip64 data.b64
```

### Example 2: Pipe base64 content through stdin

```bash
echo "ykjNyclXKM8vyklRBAQAAP//" | ./unzip64
```

This will output: `hello world!`

### Example 3: Chain with other commands

```bash
curl -s https://example.com/compressed.b64 | ./unzip64 > output.txt
```

## How it works

1. **Input**: The tool reads base64-encoded content from either:
   - A file specified as a command line argument
   - Standard input (stdin) if no file is specified

2. **Decode**: The base64 content is decoded to binary data

3. **Decompress**: The binary data is decompressed using flate

4. **Output**: The decompressed content is written to standard output (stdout)

## Error Handling

The tool will exit with an error message if:
- The specified file cannot be opened
- The input cannot be decoded as valid base64
- The decoded data cannot be decompressed as flate
- Any I/O operation fails

## Development

### Building

Use the provided Makefile for common development tasks:

```bash
make build      # Build the binary
make test       # Run tests
make build-all  # Build for all platforms
make demo       # Run demo with test data
make clean      # Clean build artifacts
make help       # Show all available targets
```

### Testing

Run the test suite:

```bash
make test
```

Or use Go directly:

```bash
go test -v ./...
```

### Continuous Integration

This project uses GitHub Actions for:
- **Continuous Integration**: Tests are run on every push and pull request
- **Nightly Releases**: Automated builds and releases every night at 2:00 AM UTC
- **Multi-platform Builds**: Binaries are built for Linux, macOS, and Windows (x86_64 and ARM64)

## Requirements

- Go 1.21 or later

## License

This project is licensed under the MIT License - see the [LICENSE](#license) file for details.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License

Copyright (c) 2025

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
