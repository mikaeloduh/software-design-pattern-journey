package entity

type Fire struct{}

func NewFire() *Fire {
	return &Fire{}
}

func (f *Fire) String() string {
	return "F"
}
