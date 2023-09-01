package entity

type Monster struct {
	MaxHp    int
	Hp       int
	Speed    int
	State    IState
	Position *Position
}

func (m *Monster) OnRoundStart() {
}

func (m *Monster) AfterRoundStart() {
}

func (m *Monster) OnRoundEnd() {
}

func (m *Monster) isRoundEnd() bool {
	return false
}

func (m *Monster) SetState(s IState) {
	m.State = s
}

func (m *Monster) SetSpeed(speed int) {
	m.Speed = speed
}

func (m *Monster) SetPosition(p *Position) {
	m.Position = p
}

func (m *Monster) TakeDamage(damage int) int {
	m.Hp -= damage

	return m.Hp
}
