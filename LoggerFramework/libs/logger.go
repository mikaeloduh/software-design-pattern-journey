package libs

// ILogger interface
type ILogger interface {
	Log(level Level, text string)
	GetExporter() IExporter
	GetLayout() ILayout
	GetLevelThreshold() Level
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

	if parent != nil {
		var levelThresholdValue Level
		if levelThreshold != UNDEFINED {
			levelThresholdValue = levelThreshold
		} else {
			levelThresholdValue = parent.GetLevelThreshold()
		}

		var exporterInstance IExporter
		if exporter != nil {
			exporterInstance = exporter
		} else {
			exporterInstance = parent.GetExporter()
		}

		var layoutInstance ILayout
		if layout != nil {
			layoutInstance = layout
		} else {
			layoutInstance = parent.GetLayout()
		}

		return &Logger{
			Parent:         parent,
			Name:           name,
			LevelThreshold: levelThresholdValue,
			Exporter:       exporterInstance,
			Layout:         layoutInstance,
		}
	}

	return &Logger{
		Parent:         nil,
		Name:           "root",
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

func (l *Logger) GetLevelThreshold() Level {
	return l.LevelThreshold
}

func (l *Logger) GetExporter() IExporter {
	return l.Exporter
}

func (l *Logger) GetLayout() ILayout {
	return l.Layout
}
