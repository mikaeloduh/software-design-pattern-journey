package commons

import "math/rand"

// RandNonRepeatInt
func RandNonRepeatInt(min, max, count int) []int {
	if max-min+1 < count {
		return nil
	}

	numbers := make([]int, max-min+1)
	for i := min; i <= max; i++ {
		numbers[i-min] = i
	}

	for i := len(numbers) - 1; i >= len(numbers)-count; i-- {
		j := rand.Intn(i + 1)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	return numbers[len(numbers)-count:]
}

// HashSet
type HashSet map[interface{}]struct{}

func NewHashSet() HashSet {
	return make(HashSet)
}

func (hs HashSet) Add(item interface{}) {
	hs[item] = struct{}{}
}

func (hs HashSet) Remove(item interface{}) {
	delete(hs, item)
}

func (hs HashSet) Contains(item interface{}) bool {
	_, found := hs[item]
	return found
}

func (hs HashSet) Size() int {
	return len(hs)
}

func (hs HashSet) Items() []interface{} {
	items := make([]interface{}, 0, hs.Size())
	for item := range hs {
		items = append(items, item)
	}
	return items
}

func RandBool() bool {
	return rand.Intn(2) == 1
}
