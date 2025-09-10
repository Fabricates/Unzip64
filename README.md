# Flate

A simple Go command line tool that compresses or decompresses data using flate/deflate compression from a f1. **Input**: The tool reads data from either:
   - A file specified as a command line argument
   - Standard input (stdin) if no file is specified

2. **Process**: The data is processed using:
   - **Compression** (default): Compress input using raw deflate or zlib format (-z flag)
   - **Decompression** (-d flag): Decompress using raw deflate or zlib format (-z flag)

3. **Output**: The processed content is written to standard output (stdout)din.

## Features

- Compress data using raw deflate or zlib format
- Decompress flate (zlib) or raw deflate compressed data
- Read from a specified file or from stdin
- Output to stdout
- Support for both zlib format (-z flag) and raw deflate format (default)

## Installation

### Download pre-built binaries

Pre-built binaries are available from the [releases page](https://github.com/fabricates/flate/releases). Nightly releases are automatically built and published.

### Build from source

```bash
git clone https://github.com/fabricates/flate.git
cd flate
make build
```

Or without Make:

```bash
go build -o flate .
```

### Install using go install

```bash
go install github.com/fabricates/flate@latest
```

## Usage

### Compress a file (default behavior)

```bash
./flate input.txt > compressed.bin
```

### Compress using zlib format

```bash
./flate -z input.txt > compressed.zlib
```

### Decompress a file

```bash
./flate -d compressed.bin
```

### Decompress zlib format

```bash
./flate -d -z compressed.zlib
```

### Read from stdin

```bash
cat input.txt | ./flate > compressed.bin
```

### Show help

```bash
./flate -h
```

## Examples

### Example 1: Compress and decompress text

```bash
# Compress text with raw deflate
echo "hello world!" | ./flate > compressed.bin

# Decompress it back
./flate -d compressed.bin
```

### Example 2: Compress with zlib format

```bash
# Compress with zlib
echo "hello world!" | ./flate -z > compressed.zlib

# Decompress zlib
./flate -d -z compressed.zlib
```

### Example 3: Chain with other commands

```bash
curl -s https://example.com/data.txt | ./flate -z | ./flate -d -z
```

### Example 4: Compress files

```bash
# Compress a file
./flate largefile.txt > largefile.deflate

# Decompress it
./flate -d largefile.deflate > restored.txt
```

## How it works

1. **Input**: The tool reads compressed binary data from either:
   - A file specified as a command line argument
   - Standard input (stdin) if no file is specified

2. **Decompress**: The binary data is decompressed using:
   - zlib format (flate with headers) - default
   - Raw deflate format - when using the `-d` flag

3. **Output**: The decompressed content is written to standard output (stdout)

## Error Handling

The tool will exit with an error message if:
- The specified file cannot be opened
- The input data cannot be processed (compressed/decompressed) as flate/deflate
- The compression format doesn't match the selected mode (zlib vs raw deflate) during decompression
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
