package entity

type TagEvent struct {
	TaggedBy Taggable
	TaggedTo Taggable
	Message  Message
}

func (e *TagEvent) GetEventName() string {
	return "TagEvent"
}

func (e *TagEvent) GetEventData() (Taggable, Taggable, Message) {
	return e.TaggedBy, e.TaggedTo, e.Message
}
