package utils

// HashSet
type HashSet map[interface{}]struct{}

func NewHashSet() HashSet {
	return make(HashSet)
}

func (hs *HashSet) Add(item interface{}) {
	(*hs)[item] = struct{}{}
}

func (hs *HashSet) Remove(item interface{}) {
	delete(*hs, item)
}

func (hs *HashSet) Contains(item interface{}) bool {
	_, found := (*hs)[item]
	return found
}

func (hs *HashSet) Size() int {
	return len(*hs)
}

func (hs *HashSet) Items() []interface{} {
	items := make([]interface{}, 0, hs.Size())
	for item := range *hs {
		items = append(items, item)
	}
	return items
}
