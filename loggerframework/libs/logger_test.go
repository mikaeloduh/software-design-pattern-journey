package libs

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestLogger_Log(t *testing.T) {
	t.Run("test standard layout", func(t *testing.T) {
		var writer bytes.Buffer

		consoleExporter := FakeNewConsoleExporter(&writer)
		standardLayout := NewStandardLayout()
		root := NewLogger(nil, "root", DEBUG, consoleExporter, standardLayout)

		root.Log(DEBUG, "hello world")

		assert.Regexp(t, `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3} \|-DEBUG root - hello world`, writer.String())
		writer.Reset()
	})

	t.Run("test log threshold level at DEBUG, should only proceed DEBUG, WARN, ERROR", func(t *testing.T) {
		var writer bytes.Buffer

		consoleExporter := FakeNewConsoleExporter(&writer)
		standardLayout := NewStandardLayout()
		root := NewLogger(nil, "root", DEBUG, consoleExporter, standardLayout)

		root.Log(DEBUG, "debug message")
		assert.Regexp(t, `\|\-DEBUG`, writer.String())
		writer.Reset()

		root.Log(WARN, "warn message")
		assert.Regexp(t, `\|\-WARN`, writer.String())
		writer.Reset()

		root.Log(ERROR, "error message")
		assert.Regexp(t, `\|\-ERROR`, writer.String())
		writer.Reset()

		root.Log(INFO, "info message")
		assert.Empty(t, writer.String())
		writer.Reset()

		root.Log(TRACE, "trace message")
		assert.Empty(t, writer.String())
		writer.Reset()
	})
}

func FakeNewConsoleExporter(w io.Writer) *ConsoleExporter {
	return &ConsoleExporter{Writer: w}
}
