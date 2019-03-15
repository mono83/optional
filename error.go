package optional

// Error is optional container for errors
type Error struct {
	value error
}

// OfError creates new optional error holder containing provided value
func OfError(err error) Error {
	return Error{value: err}
}

// Panic method raises a panic if optional contains value
func (e Error) Panic() {
	if e.IsPresent() {
		panic(e.value)
	}
}
