package entity

type Taggable interface {
	Tag(event TagEvent)
	Id() string
}
