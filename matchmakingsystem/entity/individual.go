package entity

type Individual struct {
	Id     int
	Gender Gender
	Age    int
	Intro  string
	Habits []string
	Coord  Coord
}
