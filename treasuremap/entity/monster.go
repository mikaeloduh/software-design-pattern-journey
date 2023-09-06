package entity

type Monster struct {
	MaxHp    Hp
	Hp       Hp
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

func (m *Monster) Heal(hp Hp) {
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

func (m *Monster) GetHp() Hp {
	return m.Hp
}

func (m *Monster) GetMaxHp() Hp {
	return m.MaxHp
}

func (m *Monster) GetPosition() *Position {
	return m.Position
}

func (m *Monster) TakeDamage(damage Damage) Hp {
	m.Hp -= Hp(damage)

	return m.Hp
}

func (m *Monster) Attack() {
	m.Position.Game.Attack(func(worldMap [10][10]*Position) (damageArea AttackDamageArea) {
		// TODO: implement it
		return
	})
}
