package entity

type TurnMove struct {
	player   IBigTwoPlayer
	selected int
}

func NewTurnMove(player IBigTwoPlayer, selected int) *TurnMove {
	return &TurnMove{player: player, selected: selected}
}

func (t *TurnMove) DryRun() BigTwoCard {
	return t.player.GetHand()[t.selected]
}

func (t *TurnMove) Play() BigTwoCard {
	return t.player.RemoveCard(t.selected)
}
