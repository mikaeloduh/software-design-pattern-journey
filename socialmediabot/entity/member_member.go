package entity

// Member
type Member struct {
	id string
}

func NewMember(id string) *Member {
	return &Member{id: id}
}

func (b *Member) Tag(event TagEvent) {
	panic("unimplemented")
}

func (b *Member) Id() string {
	return b.id
}
