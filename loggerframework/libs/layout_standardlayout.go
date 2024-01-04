package libs

import (
	"fmt"
	"time"
)

// StandardLayout
type StandardLayout struct{}

func NewStandardLayout() *StandardLayout {
	return &StandardLayout{}
}

func (l *StandardLayout) Print(loggerName string, message Message) string {
	// yyyy-MM-dd HH:mm:ss.SSS |-LEVEL LOGGER_NAME - CONTENT
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05.000")
	return fmt.Sprintf("%s |-%s %s - %s", formattedTime, message.Level.String(), loggerName, message.Content)
}
