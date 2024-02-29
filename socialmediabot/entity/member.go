package entity

type Member struct {
	Id string
}

func NewMember(id string) *Member {
	return &Member{Id: id}
}
