package entity

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

func (b *Member) Id() string {
	return b.id
}

func (b *Member) SetUpdater(f IUpdater) {
	b.updater = f
}

// IUpdater
type IUpdater interface {
	Do()
}
