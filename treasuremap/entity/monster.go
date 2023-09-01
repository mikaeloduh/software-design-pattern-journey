package entity

type Monster struct {
	position *Position
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
}

func (m *Monster) SetSpeed(speed int) {
}

func (m *Monster) SetPosition(p *Position) {
}
