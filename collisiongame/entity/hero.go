package entity

type Hero struct {
	Sprite
	Hp int
}

func (h *Hero) String() string {
	return "H"
}

func (h *Hero) SetHp(n int) {
	h.Hp += n
}

func NewHero() *Hero {
	return &Hero{Hp: 30}
}
