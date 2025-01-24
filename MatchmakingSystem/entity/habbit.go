package entity

type Habit string
type Habits []Habit

func (h Habits) CountIntersection(others Habits) int {
	mySet := make(map[Habit]bool)
	for _, v := range h {
		mySet[v] = true
	}
	othersSet := make(map[Habit]bool)
	for _, v := range others {
		othersSet[v] = true
	}
	count := 0
	for k, _ := range othersSet {
		if mySet[k] {
			count++
		}
	}

	return count
}
