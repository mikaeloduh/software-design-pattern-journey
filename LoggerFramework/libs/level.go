package libs

type Level int

func (l Level) String() string {
	return map[Level]string{
		UNDEFINED: "UNDEFINED",
		TRACE:     "TRACE",
		INFO:      "INFO",
		DEBUG:     "DEBUG",
		WARN:      "WARN",
		ERROR:     "ERROR",
	}[l]
}

const (
	UNDEFINED Level = 0
	TRACE     Level = 1
	INFO      Level = 2
	DEBUG     Level = 3
	WARN      Level = 4
	ERROR     Level = 5
)
