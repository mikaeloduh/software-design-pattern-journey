package service

// IMember
type IMember interface {
	Tag(event TagEvent)
	Id() string
	Role() Role
}
