package optional

// Int represents optional integer value
type Int struct {
	value    int
	presents bool
}

func (i *Int) set(value int) {
	i.value = value
	i.presents = true
}

// OfInt creates new optional integer containing provided value
func OfInt(i int) Int {
	return Int{value: i, presents: true}
}

// FilterZero applies zero value filtering
// This method will return empty optional if value inside optional is zero or missing
func (i Int) FilterZero() Int {
	if !i.IsPresent() || i.value == 0 {
		return Int{}
	}

	return i
}
