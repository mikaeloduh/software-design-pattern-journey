package service

type Taggable interface {
	Tag(event TagEvent)
	Id() string
}
