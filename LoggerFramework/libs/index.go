package libs

type LoggerFramework struct {
	loggers map[string]*Logger
}

func NewLoggerFramework() *LoggerFramework {
	return &LoggerFramework{loggers: make(map[string]*Logger)}
}

func (f *LoggerFramework) DeclareLoggers(loggers ...*Logger) {
	for _, logger := range loggers {
		f.loggers[logger.Name] = logger
	}
}

func (f *LoggerFramework) GetLoggers(name string) *Logger {
	return f.loggers[name]
}
