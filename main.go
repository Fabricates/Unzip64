package main

import (
	"compress/flate"
	"compress/zlib"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var decompress bool
	var useZlib bool
	var showHelp bool

	flag.BoolVar(&decompress, "d", false, "Decompress data (default: compress)")
	flag.BoolVar(&useZlib, "z", false, "Use zlib format (default: raw deflate)")
	flag.BoolVar(&showHelp, "h", false, "Show help message")
	flag.BoolVar(&showHelp, "help", false, "Show help message")
	flag.Parse()

	if showHelp {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [file]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Compress or decompress flate/deflate data from file or stdin\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nIf no file is specified, input is read from stdin.\n")
		fmt.Fprintf(os.Stderr, "By default, raw deflate format is used. Use -z for zlib format.\n")
		os.Exit(0)
	}

	var input io.Reader
	args := flag.Args()

	// Check if a file argument is provided
	if len(args) > 0 {
		filename := args[0]
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Error opening file %s: %v", filename, err)
		}
		defer file.Close()
		input = file
	} else {
		// Read from stdin
		input = os.Stdin
	}

	if decompress {
		// Decompression mode
		var reader io.ReadCloser
		var err error

		if useZlib {
			// Create a zlib reader
			reader, err = zlib.NewReader(input)
			if err != nil {
				log.Fatalf("Error creating zlib reader: %v", err)
			}
		} else {
			// Create a raw deflate reader
			reader = flate.NewReader(input)
		}
		defer reader.Close()

		// Copy decompressed data to stdout
		_, err = io.Copy(os.Stdout, reader)
		if err != nil {
			log.Fatalf("Error decompressing data: %v", err)
		}
	} else {
		// Compression mode (default)
		var writer io.WriteCloser
		var err error

		if useZlib {
			// Create a zlib writer
			writer = zlib.NewWriter(os.Stdout)
		} else {
			// Create a raw deflate writer
			writer, err = flate.NewWriter(os.Stdout, flate.DefaultCompression)
			if err != nil {
				log.Fatalf("Error creating deflate writer: %v", err)
			}
		}
		defer writer.Close()

		// Copy input data to compressed output
		_, err = io.Copy(writer, input)
		if err != nil {
			log.Fatalf("Error compressing data: %v", err)
		}
	}
}
