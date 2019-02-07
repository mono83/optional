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
