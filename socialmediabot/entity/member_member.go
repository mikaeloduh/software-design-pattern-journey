package entity

type Role string

var (
	ADMIN Role = "ADMIN"
	USER  Role = "USER"
)

// Member
type Member struct {
	id   string
	role Role
}

func NewMember(id string, role Role) *Member {
	return &Member{id: id, role: role}
}

func (b *Member) Tag(event TagEvent) {
}

func (b *Member) Id() string {
	return b.id
}
