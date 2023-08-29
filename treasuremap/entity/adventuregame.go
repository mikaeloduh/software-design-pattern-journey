package entity

type Round int

type AdventureGame struct {
	WorldMap  [10][10]*Position
	character *Character
	round     Round
}

func NewAdventureGame(character *Character) *AdventureGame {
	return &AdventureGame{
		character: character,
	}
}

func (g *AdventureGame) StartRound() {
	g.round++
	g.character.OnRoundStart()
}
