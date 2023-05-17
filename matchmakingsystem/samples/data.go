package samples

import "matchmakingsystem/entity"

var P1 = entity.Individual{
	Id:     1,
	Gender: entity.Male,
	Age:    10,
	Intro:  "Hello Intro",
	Habits: []string{"baseball", "cook", "sleep"},
	Coord: entity.Coord{
		X: 10,
		Y: 10,
	},
}

var P2 = entity.Individual{
	Id:     2,
	Gender: entity.Female,
	Age:    20,
	Intro:  "Hi there",
	Habits: []string{"music", "sleep", "travel"},
	Coord: entity.Coord{
		X: 5,
		Y: 5,
	},
}

var P3 = entity.Individual{
	Id:     3,
	Gender: entity.Female,
	Age:    30,
	Intro:  "Hey",
	Habits: []string{"baseball", "sports", "reading", "sleep"},
	Coord: entity.Coord{
		X: 15,
		Y: 15,
	},
}
