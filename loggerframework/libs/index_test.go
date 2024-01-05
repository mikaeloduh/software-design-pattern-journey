package libs

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestLoggerFramework_IntegrationTest(t *testing.T) {
	t.Run("test hierarchy logger should output to the correct exporter and logs message", func(t *testing.T) {
		// io writer for testing
		var rootWriter bytes.Buffer
		var gameWriter bytes.Buffer

		// Temporary file for testing
		gameLogTempFile := setupFileExporter(t)
		gameLogBackupTempFile := setupFileExporter(t)

		root := NewLogger(nil, "root", DEBUG, FakeNewConsoleExporter(&rootWriter), NewStandardLayout())
		gameLogger := NewLogger(root, "app.game", INFO,
			NewCompositeExporter(
				FakeNewConsoleExporter(&gameWriter),
				NewCompositeExporter(
					NewFileExporter(gameLogTempFile.Name()),
					NewFileExporter(gameLogBackupTempFile.Name()),
				),
			), nil)
		aiLogger := NewLogger(gameLogger, "app.game.ai", TRACE, nil, nil)

		loggerFramework := NewLoggerFramework()
		loggerFramework.DeclareLoggers(root, gameLogger, aiLogger)

		// Test app.game
		appGameLogger := loggerFramework.GetLoggers("app.game")
		appGameLogger.Info("log for game")

		// assert app.game console exporter
		assert.Regexp(t, `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3} \|-INFO app.game - log for game`, gameWriter.String())
		gameWriter.Reset()

		// assert app.game file exporter
		assertFileContent(t, gameLogTempFile, `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3} \|-INFO app.game - log for game`)
		assertFileContent(t, gameLogBackupTempFile, `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3} \|-INFO app.game - log for game`)

		// Test app.game.ai
		appGameAiLogger := loggerFramework.GetLoggers("app.game.ai")
		appGameAiLogger.Trace("log for ai")

		// assert app.game console exporter
		assert.Regexp(t, `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3} \|-TRACE app.game.ai - log for ai`, gameWriter.String())
		gameWriter.Reset()
	})
}
