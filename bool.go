package optional

// Bool represents optional boolean value
type Bool struct {
	value    bool
	presents bool
}

func (b *Bool) set(value bool) {
	b.value = value
	b.presents = true
}

// OfBool creates new optional boolean containing provided value
func OfBool(b bool) Bool {
	return Bool{value: b, presents: true}
}

// OrFalse returns false for empty optional value
// For other cases it will return optional value
func (b Bool) OrFalse() bool {
	if b.IsPresent() {
		return b.value
	}

	return false
}
