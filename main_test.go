package main

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"io"
	"os"
	"os/exec"
	"testing"
)

func TestZlibCompression(t *testing.T) {
	// Build the binary for testing
	cmd := exec.Command("go", "build", "-o", "flate_test", ".")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("flate_test")

	// Test data
	testString := "hello world!"

	// Test compression with zlib format
	cmd = exec.Command("./flate_test", "-z")
	cmd.Stdin = bytes.NewReader([]byte(testString))

	compressedOutput, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to run flate compression: %v", err)
	}

	// Now decompress to verify
	reader, err := zlib.NewReader(bytes.NewReader(compressedOutput))
	if err != nil {
		t.Fatalf("Failed to create zlib reader: %v", err)
	}
	defer reader.Close()

	var decompressed bytes.Buffer
	_, err = io.Copy(&decompressed, reader)
	if err != nil {
		t.Fatalf("Failed to decompress: %v", err)
	}

	result := decompressed.String()
	if result != testString {
		t.Errorf("Expected %q, got %q", testString, result)
	}
}

func TestZlibDecompression(t *testing.T) {
	// Build the binary for testing
	cmd := exec.Command("go", "build", "-o", "flate_test", ".")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("flate_test")

	// Test data
	testString := "hello world!"

	// Compress with zlib (flate with headers)
	var buf bytes.Buffer
	writer := zlib.NewWriter(&buf)

	_, err = writer.Write([]byte(testString))
	if err != nil {
		t.Fatalf("Failed to write to zlib writer: %v", err)
	}

	err = writer.Close()
	if err != nil {
		t.Fatalf("Failed to close zlib writer: %v", err)
	}

	// Test decompression with -d -z flags
	cmd = exec.Command("./flate_test", "-d", "-z")
	cmd.Stdin = bytes.NewReader(buf.Bytes())

	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to run flate decompression: %v", err)
	}

	result := string(output)
	if result != testString {
		t.Errorf("Expected %q, got %q", testString, result)
	}
}

func TestDeflateCompression(t *testing.T) {
	// Build the binary for testing
	cmd := exec.Command("go", "build", "-o", "flate_test", ".")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("flate_test")

	// Test data
	testString := "hello world!"

	// Test compression with raw deflate (default)
	cmd = exec.Command("./flate_test")
	cmd.Stdin = bytes.NewReader([]byte(testString))

	compressedOutput, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to run flate compression: %v", err)
	}

	// Now decompress to verify
	reader := flate.NewReader(bytes.NewReader(compressedOutput))
	defer reader.Close()

	var decompressed bytes.Buffer
	_, err = io.Copy(&decompressed, reader)
	if err != nil {
		t.Fatalf("Failed to decompress: %v", err)
	}

	result := decompressed.String()
	if result != testString {
		t.Errorf("Expected %q, got %q", testString, result)
	}
}

func TestDeflateDecompression(t *testing.T) {
	// Build the binary for testing
	cmd := exec.Command("go", "build", "-o", "flate_test", ".")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("flate_test")

	// Test data
	testString := "hello world!"

	// Compress with raw deflate
	var buf bytes.Buffer
	writer, err := flate.NewWriter(&buf, flate.DefaultCompression)
	if err != nil {
		t.Fatalf("Failed to create deflate writer: %v", err)
	}

	_, err = writer.Write([]byte(testString))
	if err != nil {
		t.Fatalf("Failed to write to deflate writer: %v", err)
	}

	err = writer.Close()
	if err != nil {
		t.Fatalf("Failed to close deflate writer: %v", err)
	}

	// Test decompression with -d flag
	cmd = exec.Command("./flate_test", "-d")
	cmd.Stdin = bytes.NewReader(buf.Bytes())

	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to run flate decompression: %v", err)
	}

	result := string(output)
	if result != testString {
		t.Errorf("Expected %q, got %q", testString, result)
	}
}
