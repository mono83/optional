package optional

var emptyInt = Int{}

// Int represents optional integer value
type Int struct {
	*int
}

// OfInt creates new optional integer containing provided value
func OfInt(i int) Int {
	return Int{int: &i}
}

// FilterZero applies zero value filtering
// This method will return empty optional if value inside optional is zero or missing
func (i Int) FilterZero() Int {
	if !i.IsPresent() || *i.int == 0 {
		return emptyInt
	}

	return i
}
