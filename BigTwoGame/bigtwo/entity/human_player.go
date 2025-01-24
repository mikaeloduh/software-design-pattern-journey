package entity

import (
	"bufio"
	"fmt"
	"io"
)

type HumanPlayer struct {
	Name        string
	Hand        []BigTwoCard
	ActionCards []BigTwoCard
	Reader      io.Reader
	Writer      io.Writer
}

func (p *HumanPlayer) Rename() {
	fmt.Fprintln(p.Writer, "Input your name")

	reader := bufio.NewScanner(p.Reader)
	reader.Scan()
	line := reader.Text()
	p.Name = line
}

func (p *HumanPlayer) SetCard(card BigTwoCard) {
	//TODO implement me
	panic("implement me")
}

func (p *HumanPlayer) RemoveCard(idx int) BigTwoCard {
	//TODO implement me
	panic("implement me")
}

func (p *HumanPlayer) GetName() string {
	return p.Name
}

func (p *HumanPlayer) GetHand() []BigTwoCard {
	//TODO implement me
	panic("implement me")
}

func (p *HumanPlayer) AddPoint() {
	//TODO implement me
	panic("implement me")
}

func (p *HumanPlayer) TakeTurnMove() *TurnMove {
	//TODO implement me
	panic("implement me")
}

func (p *HumanPlayer) SetActionCard(card BigTwoCard) {
	p.ActionCards = append(p.ActionCards, card)
}
