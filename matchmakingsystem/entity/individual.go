package entity

type Gender int8

const (
	male   Gender = 1
	female Gender = 0
)

type Coord struct {
	x float32
	y float32
}

type Individual struct {
	id     int
	gender Gender
	age    int
	intro  string
	habits []string
	coord  Coord
}

var p1 Individual = Individual{
	id:     1,
	gender: male,
	age:    10,
	intro:  "Hello intro",
	habits: []string{"baseball", "cook", "sleep"},
	coord: Coord{
		x: 10,
		y: 10,
	},
}
