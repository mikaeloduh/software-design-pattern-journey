package libs

type ILayout interface {
	Print(loggerName string, message Message) string
}
