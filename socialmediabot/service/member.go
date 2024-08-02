package service

type IMember interface {
	Tag(event TagEvent)
	Id() string
	Role() Role
}
