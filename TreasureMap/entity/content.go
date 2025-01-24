package entity

// IContent
type IContent interface {
	TouchState(object IStatefulMapObject)
}

type SuperStar struct {
}

func (c SuperStar) TouchState(object IStatefulMapObject) {
	object.SetState(NewInvincibleState(object))
}

type Poison struct {
}

func (p Poison) TouchState(object IStatefulMapObject) {
	//TODO implement me
	panic("implement me")
}

type AcceleratingPotion struct {
}

func (a AcceleratingPotion) TouchState(object IStatefulMapObject) {
	//TODO implement me
	panic("implement me")
}

type HealingPotion struct {
}

func (h HealingPotion) TouchState(object IStatefulMapObject) {
	//TODO implement me
	panic("implement me")
}

type DevilFruit struct {
}

func (d DevilFruit) TouchState(object IStatefulMapObject) {
	//TODO implement me
	panic("implement me")
}

type KingsRock struct {
}

func (k KingsRock) TouchState(object IStatefulMapObject) {
	//TODO implement me
	panic("implement me")
}

type DokodemoDoor struct {
}

func (d DokodemoDoor) TouchState(object IStatefulMapObject) {
	//TODO implement me
	panic("implement me")
}
