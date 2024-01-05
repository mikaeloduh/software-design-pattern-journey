package libs

import (
	"fmt"
	"os"
)

// FileExporter
type FileExporter struct {
	FilePath string
}

func NewFileExporter(filePath string) *FileExporter {
	return &FileExporter{FilePath: filePath}
}

func (e *FileExporter) Write(s string) {
	// Open the file in append mode
	file, err := os.OpenFile(e.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Write the string with a newline to the file
	_, err = fmt.Fprintln(file, s)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
