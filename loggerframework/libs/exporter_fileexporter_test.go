package libs

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExporter_Write(t *testing.T) {
	// Temporary file for testing
	tempFile := setupFileExporter(t)
	defer os.Remove(tempFile.Name())

	exporter := NewFileExporter(tempFile.Name())

	// Test exporter.Write
	testString := "Hello, world!"
	exporter.Write(testString)

	// Assert file content
	assertFileContent(t, tempFile, testString)
}

// Helper function to set a temporary file
func setupFileExporter(t *testing.T) *os.File {
	tempFile, err := os.CreateTemp("", "fileexporter")
	assert.NoError(t, err)

	return tempFile
}

// Helper function to assert file content
func assertFileContent(t *testing.T, file *os.File, expectedContent string) {
	file.Seek(0, 0) // Reset the file pointer to the beginning of the file
	scanner := bufio.NewScanner(file)
	var got string
	if scanner.Scan() {
		got = scanner.Text()
	}
	assert.NoError(t, scanner.Err())
	assert.Regexpf(t, expectedContent, strings.TrimSpace(got), "The written string should match the expected content")
}
