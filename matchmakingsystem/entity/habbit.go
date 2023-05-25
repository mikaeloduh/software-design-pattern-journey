package entity

type Habit string
type Habits []Habit

func (h *Habits) CountIntersection(other Habits) int {
	count := 0
	for _, x := range *h {
		for _, y := range other {
			if x == y {
				count++
				break
			}
		}
	}
	return count
}
