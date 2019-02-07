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
