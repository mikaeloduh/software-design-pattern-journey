package libs

import (
	"fmt"
	"io"
	"os"
)

type IExporter interface {
	Write(s string)
}

// ConsoleExporter
type ConsoleExporter struct {
	Writer io.Writer
}

func NewConsoleExporter() *ConsoleExporter {
	return &ConsoleExporter{Writer: os.Stdout}
}

func (e *ConsoleExporter) Write(s string) {
	_, _ = fmt.Fprintf(e.Writer, "%s", s)
}

// CompositeExporter
type CompositeExporter struct {
	Children []IExporter
}

func NewCompositeExporter(children ...IExporter) *CompositeExporter {
	return &CompositeExporter{Children: children}
}

func (e *CompositeExporter) Write(s string) {
	for _, child := range e.Children {
		child.Write(s)
	}
}
