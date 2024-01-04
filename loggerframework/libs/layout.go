package libs

type ILayout interface {
	Print(message Message) string
}

// StandardLayout
type StandardLayout struct{}

func NewStandardLayout() *StandardLayout {
	return &StandardLayout{}
}

func (l *StandardLayout) Print(message Message) string {
	return message.Content
}
