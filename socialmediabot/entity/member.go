package entity

// IMember
type IMember interface {
	Tag(event TagEvent)
	Id() string
}
