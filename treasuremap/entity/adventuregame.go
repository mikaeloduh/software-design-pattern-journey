package entity

type AdventureGame struct {
	character *Character
	round     int
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
