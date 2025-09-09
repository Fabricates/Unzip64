package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	// Build the binary for testing
	cmd := exec.Command("go", "build", "-o", "unzip64_test", ".")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("unzip64_test")

	// Test data
	testString := "hello world!"

	// Compress with flate
	var buf bytes.Buffer
	writer, err := flate.NewWriter(&buf, flate.DefaultCompression)
	if err != nil {
		t.Fatalf("Failed to create flate writer: %v", err)
	}

	_, err = writer.Write([]byte(testString))
	if err != nil {
		t.Fatalf("Failed to write to flate writer: %v", err)
	}

	err = writer.Close()
	if err != nil {
		t.Fatalf("Failed to close flate writer: %v", err)
	}

	// Encode to base64
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Test the binary
	cmd = exec.Command("./unzip64_test")
	cmd.Stdin = strings.NewReader(encoded)

	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to run unzip64: %v", err)
	}

	result := string(output)
	if result != testString {
		t.Errorf("Expected %q, got %q", testString, result)
	}
}

func TestBase64Decode(t *testing.T) {
	// Test with known good data
	testData := "ykjNyclXKM8vyklRBAQAAP//"
	expected := "hello world!"

	// Build the binary for testing
	cmd := exec.Command("go", "build", "-o", "unzip64_test", ".")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("unzip64_test")

	// Test the binary
	cmd = exec.Command("./unzip64_test")
	cmd.Stdin = strings.NewReader(testData)

	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to run unzip64: %v", err)
	}

	result := string(output)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
