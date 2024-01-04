package libs

import (
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
