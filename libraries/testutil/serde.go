package testutil

import (
	"io"
	"os"
	"testing"
)

func CreateBody(t *testing.T, path string) io.Reader {
	t.Helper()

	reader, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open file, %v", err)
	}

	return reader
}

func ReadBody(t *testing.T, reader io.Reader) []byte {
	t.Helper()

	bytes, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("failed to decode input, %v", err)
	}

	return bytes
}

func ReadFile(t *testing.T, path string) []byte {
	t.Helper()

	bytes, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read file from path %v, %v", path, err)
	}

	return bytes
}
