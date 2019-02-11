package optional

var emptyBool = Bool{}

// Bool represents optional boolean value
type Bool struct {
	*bool
}

// OfBool creates new optional boolean containing provided value
func OfBool(b bool) Bool {
	return Bool{bool: &b}
}

// OrFalse returns false for empty optional value
// For other cases it will return optional value
func (b Bool) OrFalse() bool {
	if b.IsPresent() {
		return *b.bool
	}

	return false
}
