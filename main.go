package main

import (
	"bufio"
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
	"log"
	"os"
)

func main() {
	var input io.Reader

	// Check if a file argument is provided
	if len(os.Args) > 1 {
		filename := os.Args[1]
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

	// Read all input
	scanner := bufio.NewScanner(input)
	var base64Content string
	for scanner.Scan() {
		base64Content += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Decode base64
	decodedData, err := base64.StdEncoding.DecodeString(base64Content)
	if err != nil {
		log.Fatalf("Error decoding base64: %v", err)
	}

	// Create a flate reader to decompress the data
	flateReader := flate.NewReader(bytes.NewReader(decodedData))
	defer flateReader.Close()

	// Copy decompressed data to stdout
	_, err = io.Copy(os.Stdout, flateReader)
	if err != nil {
		log.Fatalf("Error decompressing data: %v", err)
	}
}
