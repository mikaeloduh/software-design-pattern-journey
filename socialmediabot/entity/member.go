package entity

// IMember
type IMember interface {
	Tag(event TagEvent)
	Id() string
}

// Member
type Member struct {
	id      string
	updater IUpdater
}

func NewMember(id string) *Member {
	return &Member{id: id}
}

func (b *Member) Tag(event TagEvent) {
	b.updater.Do()
}

func (b *Member) SetUpdater(f IUpdater) {
	b.updater = f
}

func (b *Member) Id() string {
	return b.id
}

// IUpdater
type IUpdater interface {
	Do()
}
