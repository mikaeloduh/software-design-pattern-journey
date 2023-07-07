package main

type Hero struct {
	Sprite
	hp int
}

func (h *Hero) String() string {
	return "H"
}

func (h *Hero) SetHp(n int) {
	h.hp += n
}

func NewHero() *Hero {
	return &Hero{hp: 30}
}
