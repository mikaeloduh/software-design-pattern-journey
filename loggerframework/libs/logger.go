package libs

type Message struct {
	Level   Level
	Content string
}

type Level int

func (l Level) String() string {
	return map[Level]string{
		TRACE: "TRACE",
		INFO:  "INFO",
		DEBUG: "DEBUG",
		WARN:  "WARN",
		ERROR: "ERROR",
	}[l]
}

const (
	TRACE Level = 0
	INFO  Level = 1
	DEBUG Level = 2
	WARN  Level = 3
	ERROR Level = 4
)

type ILogger interface {
	Log(level Level, text string) Message
}

// Logger
type Logger struct {
	Parent         ILogger
	Name           string
	LevelThreshold Level
	Exporter       IExporter
	Layout         ILayout
}

func NewLogger(parent ILogger, name string, levelThreshold Level, exporter IExporter, layout ILayout) *Logger {
	return &Logger{
		Parent:         parent,
		Name:           name,
		LevelThreshold: levelThreshold,
		Exporter:       exporter,
		Layout:         layout,
	}
}

func (l *Logger) Log(level Level, text string) {
	if level < l.LevelThreshold {
		return
	}

	layout := l.Layout.Print(l.Name, Message{Level: level, Content: text})

	l.Exporter.Write(layout)
}

func (l *Logger) Trace(s string) {
	l.Log(TRACE, s)
}

func (l *Logger) Info(s string) {
	l.Log(INFO, s)
}

func (l *Logger) Debug(s string) {
	l.Log(DEBUG, s)
}

func (l *Logger) Warn(s string) {
	l.Log(WARN, s)
}

func (l *Logger) Error(s string) {
	l.Log(ERROR, s)
}
