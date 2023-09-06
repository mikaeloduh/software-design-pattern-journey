package entity

import (
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"treasuremap/utils"
)

type Character struct {
	Writer         io.Writer
	MaxHp          int
	Hp             int
	AttackDamage   int
	Speed          int // actions per round
	State          IState
	Position       *Position
	disableActions utils.HashSet
}

func NewCharacter() *Character {
	var c *Character
	c = &Character{
		Writer:       os.Stdout,
		MaxHp:        300,
		Hp:           300,
		AttackDamage: 999,
		Speed:        1,
		State:        NewNormalState(c),
	}
	return c
}

func (c *Character) OnRoundStart() {
	c.State.OnRoundStart()

	for !c.isRoundEnd() {
		c.AfterRoundStart()
	}

	c.OnRoundEnd()
}

func (c *Character) AfterRoundStart() {
	fmt.Fprint(c.Writer, "take action\n")

	c.Speed--
}

func (c *Character) OnRoundEnd() {
	c.Speed = 1
	c.disableActions = utils.NewHashSet()
}

func (c *Character) isRoundEnd() bool {
	return c.Speed <= 0
}

func (c *Character) TakeDamage(damage int) int {
	c.Hp -= c.State.OnTakeDamage(damage)

	return c.Hp
}

func (c *Character) Heal(health int) {
	c.Hp = int(math.Min(float64(float32(c.MaxHp)), float64(c.Hp+health)))
}

func (c *Character) MoveUp() {
	if c.disableActions.Contains("MoveUp") {
		return
	}
	c.Position.Move(c.Position.X, c.Position.Y+1, Up)
}

func (c *Character) MoveDown() {
	if c.disableActions.Contains("MoveDown") {
		return
	}
	c.Position.Move(c.Position.X, c.Position.Y-1, Down)
}

func (c *Character) MoveLeft() {
	if c.disableActions.Contains("MoveLeft") {
		return
	}
	c.Position.Move(c.Position.X-1, c.Position.Y, Left)
}

func (c *Character) MoveRight() {
	if c.disableActions.Contains("MoveRight") {
		return
	}
	c.Position.Move(c.Position.X+1, c.Position.Y, Right)
}

func (c *Character) Attack() {
	var defaultAttackStrategy = func(worldMap [10][10]*Position) (attackRange AttackRange) {
		// Attack eliminates all monsters in front of character in a straight line and stop at the first obstacle
		x := c.Position.X
		for y := c.Position.Y + 1; y < len(worldMap); y++ {
			if worldMap[y][x] != nil {
				fmt.Println(reflect.TypeOf(worldMap[y][x].Object))
				if reflect.TypeOf(worldMap[y][x].Object) == reflect.TypeOf(&Obstacle{}) {
					return
				}
			}
			attackRange[y][x] = c.AttackDamage
		}
		return
	}

	attackStrategy := c.State.OnAttack(defaultAttackStrategy)

	c.Position.Game.Attack(attackStrategy)
}

func (c *Character) SetState(s IState) {
	c.State = s
}

func (c *Character) SetSpeed(speed int) {
	c.Speed = speed
}

func (c *Character) SetPosition(p *Position) {
	c.Position = p
}

func (c *Character) GetHp() int {
	return c.Hp
}

func (c *Character) GetMaxHp() int {
	return c.MaxHp
}

func (c *Character) GetPosition() *Position {
	return c.Position
}

func (c *Character) DisableAction(action string) {
	c.disableActions.Add(action)
}
