package libs

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoggerFramework(t *testing.T) {
	t.Run("test DeclareLoggers and GetLoggers", func(t *testing.T) {
		root := NewLogger(nil, "root", DEBUG, NewConsoleExporter(), NewStandardLayout())
		gameLogger := NewLogger(root, "app.game", UNDEFINED, nil, nil)

		loggerFramework := NewLoggerFramework()
		loggerFramework.DeclareLoggers(root, gameLogger)

		appGameLogger := loggerFramework.GetLoggers("app.game")

		assert.Same(t, gameLogger, appGameLogger)
	})
}

func TestLoggerFramework_Log(t *testing.T) {
	t.Run("test GetLoggers should return correct logger and logs message", func(t *testing.T) {
		var writer bytes.Buffer

		root := NewLogger(nil, "root", DEBUG, FakeNewConsoleExporter(&writer), NewStandardLayout())
		gameLogger := NewLogger(root, "app.game", UNDEFINED, nil, nil)

		loggerFramework := NewLoggerFramework()
		loggerFramework.DeclareLoggers(root, gameLogger)

		appGameLogger := loggerFramework.GetLoggers("app.game")

		appGameLogger.Debug("log for game")

		assert.Regexp(t, `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3} \|-DEBUG app.game - log for game`, writer.String())
		writer.Reset()
	})
}
