package entity

// Speak
type Speak struct {
	Speaker IMember
	Content string
}

func NewSpeak(speaker IMember, content string) Speak {
	return Speak{Speaker: speaker, Content: content}
}
