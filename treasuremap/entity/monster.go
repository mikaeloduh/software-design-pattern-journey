package entity

type Monster struct {
	MaxHp    int
	Hp       int
	Speed    int
	State    IState
	Position *Position
}

func NewMonster() *Monster {
	var m *Monster
	m = &Monster{
		MaxHp: 10,
		Hp:    10,
		Speed: 1,
		State: NewNormalState(m),
	}
	return m
}

func (m *Monster) Heal(hp int) {
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

func (m *Monster) GetHp() int {
	return m.Hp
}

func (m *Monster) GetMaxHp() int {
	return m.MaxHp
}

func (m *Monster) GetPosition() *Position {
	return m.Position
}

func (m *Monster) TakeDamage(damage int) int {
	m.Hp -= damage

	return m.Hp
}
